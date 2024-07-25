package service

import (
	request "resetgoapi.com/go-admin-api/common/request/system"
	"resetgoapi.com/go-admin-api/global"
	models "resetgoapi.com/go-admin-api/models/system"
)

var DeptService = deptService{}

type deptService struct{}

type IDeptService interface {
	List() ([]*models.SysDept, error)
	Page(request *request.DeptListRequest) ([]*models.SysDept, error)
	Create(request *request.CreateDeptRequest) error
	Update(request *request.UpdateDeptRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysDept, error)
	Find() ([]*models.SysDept, error)
}

func (d deptService) List() (depts []*models.SysDept, err error) {
	err = global.GORM.Find(&depts).Error
	return
}

func (d deptService) Page(request *request.DeptListRequest) (depts []*models.SysDept, err error) {
	orm := global.GORM
	if request.Status != nil {
		orm = orm.Where("status = ?", request.Status)
	}
	orm = orm.Where("dept_name LIKE ?", "%"+request.DeptName+"%")
	err = orm.Find(&depts).Error
	return
}

func (d deptService) Create(request *request.CreateDeptRequest) (err error) {
	err = global.GORM.Create(&models.SysDept{
		DeptName: request.DeptName,
		ParentId: *request.ParentId,
		DeptSort: request.DeptSort,
		Remark:   request.Remark,
		Leader:   request.Leader,
		Phone:    request.Phone,
		Status:   &request.Status,
	}).Error
	return
}

func (d deptService) Update(request *request.UpdateDeptRequest) (err error) {
	var menu models.SysDept
	menu.Id = request.Id
	err = global.GORM.Model(&menu).Updates(&models.SysDept{
		ParentId: *request.ParentId,
		DeptSort: request.DeptSort,
		Remark:   request.Remark,
		Leader:   request.Leader,
		Phone:    request.Phone,
		Status:   request.Status,
		DeptName: request.DeptName,
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
