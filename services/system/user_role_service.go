package service

import (
	"resetgoapi.com/go-admin-api/global"
	models "resetgoapi.com/go-admin-api/models/system"
)

var UserRoleService = userRoleService{}

type userRoleService struct{}

type IUserRoleService interface {
	Create(request *models.SysUserRole) error
	FindRoleIdsByUserId(userId int64) ([]int64, error)
}

func (u userRoleService) Create(userRole *models.SysUserRole) (err error) {
	err = global.GORM.Create(&userRole).Error
	return
}

func (u userRoleService) FindRoleIdsByUserId(userId int64) ([]int64, error) {
	var sysUserRole []*models.SysUserRole
	if err := global.GORM.Where("user_id = ?", userId).Find(&sysUserRole).Error; err != nil {
		return nil, err
	}
	var roleIds []int64
	for _, role := range sysUserRole {
		roleIds = append(roleIds, role.RoleId)
	}
	return roleIds, nil
}
