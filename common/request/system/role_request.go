package request

import "resetgoapi.com/rest_go_api/common/request"

type RoleListRequest struct {
	request.PageRequest
}

type CreateRoleRequest struct {
	// 角色名称
	RoleName string `json:"roleName" form:"roleName"`
	// 角色标识
	RoleKey string `json:"roleKey" form:"roleKey"`
	// 角色排序
	RoleSort int64 `json:"roleSort" form:"roleSort"`
	// 启用状态 0 启用 1 禁用
	Status *int `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}

type UpdateRoleRequest struct {
	Id int64 `json:"id" form:"id"`
	// 角色名称
	RoleName string `json:"roleName" form:"roleName"`
	// 角色标识
	RoleKey string `json:"roleKey" form:"roleKey"`
	// 角色排序
	RoleSort int64 `json:"roleSort" form:"roleSort"`
	// 启用状态 0 启用 1 禁用
	Status *int `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
