package service

import (
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
)

var RoleDeptService = roleDeptService{}

type roleDeptService struct{}

type IRoleDeptService interface {
	Create(request *models.SysRoleDept) error
}

func (u roleDeptService) Create(roleDept *models.SysRoleDept) (err error) {
	err = global.GORM.Create(&roleDept).Error
	return
}
