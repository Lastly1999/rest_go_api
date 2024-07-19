package service

import (
	request "resetgoapi.com/rest_go_api/common/request/system"
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/pkg/db/scopes"
)

type RoleService struct{}

type IRoleService interface {
	Page(request *request.RoleListRequest) (*models.SysRole, int64, error)
	Create(request *request.CreateRoleRequest) error
	Update(request *request.UpdateRoleRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysRole, error)
	Find() ([]*models.SysRole, error)
}

func (r RoleService) Create(request *request.CreateRoleRequest) error {
	if err := global.GORM.Create(&models.SysRole{
		RoleName: request.RoleName,
		RoleKey:  request.RoleKey,
		RoleSort: request.RoleSort,
		Status:   request.Status,
		Remark:   request.Remark,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Update(request *request.UpdateRoleRequest) error {
	var role models.SysRole
	role.Id = request.Id
	if err := global.GORM.Updates(&role).Updates(&models.SysRole{
		RoleName: request.RoleName,
		RoleKey:  request.RoleKey,
		RoleSort: request.RoleSort,
		Status:   request.Status,
		Remark:   request.Remark,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Delete(id int64) error {
	if err := global.GORM.Where("id = ?", id).Delete(&models.SysRole{}).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) FindOne(id int64) (*models.SysRole, error) {
	var role models.SysRole
	if err := global.GORM.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r RoleService) Find() ([]*models.SysRole, error) {
	var roles []*models.SysRole
	if err := global.GORM.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r RoleService) Page(request *request.RoleListRequest) (roles *models.SysRole, total int64, err error) {
	err = global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&roles).Count(&total).Error
	return
}
