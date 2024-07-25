package models

import "resetgoapi.com/rest_go_api/models"

type SysRole struct {
	models.BaseModel
	// 角色名称
	RoleName string `json:"roleName" form:"roleName"`
	// 角色标识
	RoleKey string `json:"roleKey" form:"roleKey"`
	// 角色排序
	RoleSort int64 `json:"roleSort" form:"roleSort"`
	// 启用状态 0 启用 1 禁用
	Status int `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
