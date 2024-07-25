package response

import models "resetgoapi.com/go-admin-api/models/system"

type GetUserInfoResponse struct {
	Roles    []string `json:"roles"`
	UserId   int64    `json:"userId"`
	Username string   `json:"username"`
	RealName string   `json:"realName"`
	Avatar   string   `json:"avatar"`
	Desc     string   `json:"desc"`
}

type SysUserPageListResponse struct {
	Role []string `json:"role"`
	models.SysUser
}
