package service

import (
	request "resetgoapi.com/rest_go_api/common/request/system"
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/pkg/db/scopes"
)

var DeptService = deptService{}

type deptService struct{}

type IDeptService interface {
	Page(request *request.DeptListRequest) (*models.SysDept, int64, error)
	Create(request *request.CreateDeptRequest) error
	Update(request *request.UpdateDeptRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysDept, error)
	Find() ([]*models.SysDept, error)
}

func (d deptService) Page(request *request.DeptListRequest) (menus *models.SysDept, total int64, err error) {
	err = global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&menus).Count(&total).Error
	return
}

func (d deptService) Create(request *request.CreateDeptRequest) (err error) {
	err = global.GORM.Create(&models.SysDept{
		ParentId: request.ParentId,
		DeptSort: request.DeptSort,
		Remark:   request.Remark,
		Leader:   request.Leader,
		Phone:    request.Phone,
		Status:   request.Status,
	}).Error
	return
}

func (d deptService) Update(request *request.UpdateDeptRequest) (err error) {
	var menu models.SysDept
	menu.Id = request.Id
	err = global.GORM.Model(&menu).Updates(&models.SysDept{
		ParentId: request.ParentId,
		DeptSort: request.DeptSort,
		Remark:   request.Remark,
		Leader:   request.Leader,
		Phone:    request.Phone,
		Status:   request.Status,
	}).Error
	return
}

func (d deptService) Delete(id int64) (err error) {
	err = global.GORM.Where("id = ?", id).Delete(&models.SysDept{}).Error
	return
}

func (d deptService) FindOne(id int64) (dept *models.SysDept, err error) {
	err = global.GORM.Where("id = ?", id).First(&dept).Error
	return
}

func (d deptService) Find() (menus []*models.SysDept, err error) {
	err = global.GORM.Find(&menus).Error
	return
}
