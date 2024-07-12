package system

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
	NikeName string `json:"nike_name" form:"nike_name"`
	// 性别 0未知 1男 2 女
	Gender bool `json:"gender" form:"gender"`
	// 手机号
	Phone string `json:"phone" form:"phone"`
	// 头像
	Avatar string `json:"avatar" form:"avatar"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 删除状态 0 正常 1删除
	DelFlag int64 `json:"del_flag" form:"del_flag"`
	// 备注
	Remark string `json:"remark" form:"remark"`
	//
	DeptId int64 `json:"dept_id" form:"dept_id"`
	//
	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"autoCreateTime"`
	//
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime"`
}
