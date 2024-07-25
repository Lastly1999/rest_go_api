package service

import (
	request "resetgoapi.com/go-admin-api/common/request/system"
	"resetgoapi.com/go-admin-api/global"
	models "resetgoapi.com/go-admin-api/models/system"
	"resetgoapi.com/go-admin-api/pkg/db/scopes"
)

var RoleService = roleService{}

type roleService struct{}

type IRoleService interface {
	Page(request *request.RoleListRequest) ([]*models.SysRole, int64, error)
	Create(request *request.CreateRoleRequest) error
	Update(request *request.UpdateRoleRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysRole, error)
	Find() ([]*models.SysRole, error)
	FindRolesInIds(ids []int64) ([]*models.SysRole, error)
}

func (r roleService) FindRolesInIds(ids []int64) (roles []*models.SysRole, err error) {
	err = global.GORM.Where("id in (?)", ids).Find(&roles).Error
	return
}

func (r roleService) Create(request *request.CreateRoleRequest) error {
	if err := global.GORM.Create(&models.SysRole{
		RoleName: request.RoleName,
		RoleKey:  request.RoleKey,
		RoleSort: request.RoleSort,
		Status:   *request.Status,
		Remark:   request.Remark,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r roleService) Update(request *request.UpdateRoleRequest) error {
	var role models.SysRole
	role.Id = request.Id
	if err := global.GORM.Model(&role).Updates(&models.SysRole{
		RoleName: request.RoleName,
		RoleKey:  request.RoleKey,
		RoleSort: request.RoleSort,
		Status:   *request.Status,
		Remark:   request.Remark,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r roleService) Delete(id int64) error {
	if err := global.GORM.Where("id = ?", id).Delete(&models.SysRole{}).Error; err != nil {
		return err
	}
	return nil
}

func (r roleService) FindOne(id int64) (role *models.SysRole, err error) {
	err = global.GORM.Where("id = ?", id).First(&role).Error
	return
}

func (r roleService) Find() (roles []*models.SysRole, err error) {
	err = global.GORM.Find(&roles).Error
	return
}

func (r roleService) Page(request *request.RoleListRequest) (roles []*models.SysRole, total int64, err error) {
	err = global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&roles).Count(&total).Error
	return
}
