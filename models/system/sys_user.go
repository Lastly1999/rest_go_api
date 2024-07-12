package model

import (
	"time"
)

type SysUser struct {
	//
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	// 用户名
	Username string `json:"username" form:"username"`
	// 密码
	Password string `json:"password" form:"password"`
	// 邮箱
	Email string `json:"email" form:"email"`
	// 用户昵称
	NikeName string `json:"nikeName" form:"nikeName"`
	// 性别 0未知 1男 2 女
	Gender int64 `json:"gender" form:"gender"`
	// 手机号
	Phone string `json:"phone" form:"phone"`
	// 头像
	Avatar string `json:"avatar" form:"avatar"`
	// 启用状态 0 启用 1 禁用
	Status int64 `json:"status" form:"status"`
	// 删除状态 0 正常 1删除
	DelFlag int64 `json:"delFlag" form:"delFlag"`
	// 备注
	Remark string `json:"remark" form:"remark"`
	//
	DeptId int64 `json:"deptId" form:"deptId"`
	//
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"autoCreateTime"`
	//
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"autoUpdateTime"`
}