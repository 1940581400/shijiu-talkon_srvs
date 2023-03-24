package model

import "time"

type BaseModel struct {
	Id         int32     `gorm:"primaryKey"`
	CreateTime time.Time `gorm:"type:datetime comment '创建时间'"`
	UpdateTime time.Time `gorm:"type:datetime comment '更新时间'"`
	CreateUser int32     `gorm:"type:int(32) comment '创建人'"`
	UpdateUser int32     `gorm:"type:int(32) comment '更新人'"`
	IsDeleted  int       `gorm:"type:int(2) default 0 comment '是否删除,1是,0否'"`
}

type User struct {
	BaseModel
	Email    string     `gorm:"index:idx_email;type:varchar(32) comment '邮箱'"`
	Mobile   string     `gorm:"index:idx_mobile;type:varchar(11) comment '手机号'"`
	Password string     `gorm:"type:varchar(100) not null comment '密码'"`
	NickName string     `gorm:"type:varchar(20) not null comment '昵称' "`
	Birthday *time.Time `gorm:"type:datetime comment '出生日期'"`
	Gender   int        `gorm:"type:int(2) comment '性别:0男,1女'"`
	IdCard   string     `gorm:"type:varchar(32) comment '身份证号'"`
	UserType int        `gorm:"type:int(2) default 0 comment '用户类型:0普通用户,1管理员用户'"`
}
