package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"talkon_srvs/user_srv/global"
	"talkon_srvs/user_srv/global/pwd"
	"talkon_srvs/user_srv/global/random"
	"talkon_srvs/user_srv/global/zero"
	"talkon_srvs/user_srv/model"
	"talkon_srvs/user_srv/proto"
	"talkon_srvs/user_srv/utils"
	"time"
)

type UserService struct {
	proto.UnimplementedUserServer
}

func UserModelToResp(user model.User) *proto.UserInfoResp {
	userInfo := &proto.UserInfoResp{
		Id:         user.Id,
		NickName:   user.NickName,
		Mobile:     user.Mobile,
		Email:      user.Email,
		Gender:     int32(user.Gender),
		Password:   user.Password,
		IdCard:     user.IdCard,
		UserType:   int32(user.UserType),
		UpdateUser: user.UpdateUser,
		UpdateTime: uint64(user.UpdateTime.Unix()),
		CreateUser: user.CreateUser,
		CreateTime: uint64(user.CreateTime.Unix()),
		IsDeleted:  int32(user.IsDeleted),
	}
	if user.Birthday != nil {
		userInfo.Birthday = uint64(user.Birthday.Unix())
	}
	return userInfo
}

// GetUserList 分页获取用户信息
func (s *UserService) GetUserList(ctx context.Context, req *proto.PageInfoReq) (*proto.UserInfoListResp, error) {
	var userInfos []model.User
	result := global.DB.Scopes(utils.Paginate(int32(req.PageNo), int32(req.PageSize))).Find(&userInfos)
	if result.Error != nil {
		zap.L().Error("[GetUserList] 查询出错", zap.String("msg", result.Error.Error()))
		return nil, status.Error(codes.NotFound, "未查询到用户")
	}
	resp := &proto.UserInfoListResp{}
	resp.Total = result.RowsAffected
	for _, user := range userInfos {
		data := UserModelToResp(user)
		resp.Data = append(resp.Data, data)
	}
	return resp, nil
}

// GetUserById 根据id查询用户信息
func (s *UserService) GetUserById(ctx context.Context, req *proto.IdReq) (*proto.UserInfoResp, error) {
	if req.Id == zero.Number {
		return nil, status.Error(codes.NotFound, "未查询到用户")
	}
	var userInfo model.User
	result := global.DB.First(&userInfo, req.Id)
	if result.RowsAffected == 0 {
		zap.L().Error("[GetUserById] 查询出错", zap.String("msg", result.Error.Error()))
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	return UserModelToResp(userInfo), nil
}

// GetUserByMobile 通过手机号查询用户信息
func (s *UserService) GetUserByMobile(ctx context.Context, req *proto.MobileReq) (*proto.UserInfoResp, error) {
	if req.Mobile == zero.String {
		return nil, status.Error(codes.NotFound, "未查询到用户")
	}
	var userInfo model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&userInfo)
	if result.RowsAffected == 0 {
		zap.L().Error("[GetUserByMobile] 查询出错", zap.String("msg", result.Error.Error()))
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	return UserModelToResp(userInfo), nil
}

// GetUserByEmail 通过邮箱查询用户信息
func (s *UserService) GetUserByEmail(ctx context.Context, req *proto.EmailReq) (*proto.UserInfoResp, error) {
	if req.Email == zero.String {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	var userInfo model.User
	result := global.DB.Where(&model.User{Email: req.Email}).First(&userInfo)
	if result.RowsAffected == 0 {
		zap.L().Error("[GetUserByEmail] 查询出错", zap.String("msg", result.Error.Error()))
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	return UserModelToResp(userInfo), nil
}

// CreateUser 创建用户
func (s *UserService) CreateUser(ctx context.Context, req *proto.CreateUserReq) (*proto.UserInfoResp, error) {
	if req.GetMobile() == zero.String && req.GetEmail() == zero.String {
		return nil, status.Error(codes.InvalidArgument, "手机号和邮箱不能同时为空")
	}
	var existUser int64 = 0
	if req.GetMobile() != zero.String {
		global.DB.Where(&model.User{Mobile: req.Mobile}).Count(&existUser)
		if existUser > 0 {
			return nil, status.Error(codes.AlreadyExists, "手机号已被注册")
		}
	}
	if req.GetEmail() != zero.String {
		global.DB.Where(&model.User{Email: req.Email}).Count(&existUser)
		if existUser > 0 {
			return nil, status.Error(codes.AlreadyExists, "邮箱已被占用")
		}
	}
	if req.GetMobile() != zero.String {
		global.DB.Where(&model.User{Mobile: req.Mobile}).Count(&existUser)
	}
	if req.GetPassword() == zero.String {
		return nil, status.Error(codes.InvalidArgument, "密码不能为空")
	}
	now := time.Now()
	pwdSep, err := pwd.NewEncodedPwdSep(req.Password)
	if err != nil {
		return nil, err
	}
	newUser := new(model.User)
	newUser.CreateTime = now
	newUser.UpdateTime = now
	newUser.Password = pwdSep
	// 如果昵称为空，则随机生成一串字符
	if req.NickName == zero.String {
		req.NickName = random.GetStr(8)
	}
	newUser.NickName = req.NickName
	result := global.DB.Create(&newUser)
	if result.Error != nil {
		zap.L().Error("[CreateUser] 数据库操作出错", zap.String("msg", result.Error.Error()))
		return nil, status.Error(codes.Internal, "创建用户失败")
	}
	return UserModelToResp(*newUser), nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(ctx context.Context, req *proto.UpdateUserReq) (*emptypb.Empty, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在,更新失败")
	}
	birthday := time.Unix(int64(req.Birthday), 0)
	now := time.Now()
	user.NickName = req.NickName
	user.Gender = int(req.Gender)
	user.Birthday = &birthday
	user.CreateTime = now
	user.UpdateTime = now
	result = global.DB.Save(&user)
	if result.Error != nil {
		zap.L().Error("[UpdateUser] 数据库操作出错", zap.String("msg", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	return &emptypb.Empty{}, nil
}

// CheckPassword 校验密码是否正确
func (s *UserService) CheckPassword(ctx context.Context, req *proto.CheckPasswordReq) (*proto.CheckPasswordResp, error) {
	ok, err := pwd.VerifyPwdSep(req.GetEncodedPwdSep(), req.GetPassword())
	if err != nil {
		zap.L().Error("[CheckPassword] 密码校验出错", zap.String("msg", err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.CheckPasswordResp{Ok: ok}, nil
}
