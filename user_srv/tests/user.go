package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"talkon_srvs/user_srv/proto"
)

var (
	userClient proto.UserClient
	conn       *grpc.ClientConn
)

func Init() {
	var err error
	conn, err = grpc.Dial("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	resp, err := userClient.GetUserList(context.Background(), &proto.PageInfoReq{PageNo: 1, PageSize: 10})
	if err != nil {
		panic(err)
	}
	for _, data := range resp.Data {
		fmt.Println(data)
	}
}

func TestCreateUser() {
	resp, err := userClient.CreateUser(context.Background(), &proto.CreateUserReq{
		NickName: "3253",
		Password: "admin",
		Mobile:   "18173849506",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestCheckPwd() {
	resp, err := userClient.CheckPassword(context.Background(), &proto.CheckPasswordReq{
		Password:      "admin",
		EncodedPwdSep: "pbkdf2$sha512$TfBvLak1uvwEya7m$d63e2cbcdb1f037ef72c939bb609814ae71e40165c17c278b9f37fc806405ea3",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Ok)
}

func main() {
	Init()
	TestGetUserList()
	// 测试方法
	TestCheckPwd()

	err := conn.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
