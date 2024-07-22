package service

import (
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
)

var RoleMenuService = roleMenuService{}

type roleMenuService struct{}

type IRoleMenuService interface {
	Create(request *models.SysUserPost) error
}

func (u roleMenuService) Create(roleMenu *models.SysRoleMenu) (err error) {
	err = global.GORM.Create(&roleMenu).Error
	return
}
