package request

import "resetgoapi.com/go-admin-api/common/request"

type UserRequest struct {
	request.PageRequest
}

type CreateUserRequest struct {
	// 用户名
	Username string `json:"username" form:"username"`
	// 邮箱
	Email string `json:"email" form:"email"`
	// 用户昵称
	NickName string `json:"nickName" form:"nickName"`
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
	// 部门id
	DeptId int64 `json:"deptId" form:"deptId"`
}

type UpdateUserRequest struct {
	Id int64 `json:"id" form:"id"`
	// 用户名
	Username string `json:"username" form:"username"`
	// 邮箱
	Email string `json:"email" form:"email"`
	// 用户昵称
	NickName string `json:"nickName" form:"nickName"`
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
	// 部门id
	DeptId int64 `json:"deptId" form:"deptId"`
}
