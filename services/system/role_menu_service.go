package service

import (
	"resetgoapi.com/go-admin-api/global"
	models "resetgoapi.com/go-admin-api/models/system"
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
