package service

import (
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
)

var UserRoleService = userRoleService{}

type userRoleService struct{}

type IUserRoleService interface {
	Create(request *models.SysUserRole) error
}

func (u userRoleService) Create(userRole *models.SysUserRole) (err error) {
	err = global.GORM.Create(&userRole).Error
	return
}
