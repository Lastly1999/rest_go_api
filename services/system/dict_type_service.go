package service

import (
	request "resetgoapi.com/rest_go_api/common/request/system"
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/pkg/db/scopes"
)

var DictTypeService = dictTypeService{}

type dictTypeService struct{}
type IDictTypeService interface {
	Page(request *request.DictTypeListRequest) (*models.SysDictType, int64, error)
	Create(request *request.CreateDictTypeRequest) error
	Update(request *request.UpdateDictTypeRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysDictType, error)
	Find() ([]*models.SysDictType, error)
}

func (d dictTypeService) Page(request *request.DictTypeListRequest) (dictTypes *models.SysDictType, total int64, err error) {
	err = global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&dictTypes).Count(&total).Error
	return
}

func (d dictTypeService) Create(request *request.CreateDictTypeRequest) (err error) {
	err = global.GORM.Create(&models.SysDictType{
		Remark:   request.Remark,
		DictName: request.DictName,
		DictType: request.DictType,
		Status:   request.Status,
	}).Error
	return
}

func (d dictTypeService) Update(request *request.UpdateDictTypeRequest) (err error) {
	var dictType models.SysDictType
	dictType.Id = request.Id
	err = global.GORM.Model(&dictType).Updates(&models.SysDictType{
		Remark:   request.Remark,
		DictName: request.DictName,
		DictType: request.DictType,
		Status:   request.Status,
	}).Error
	return
}

func (d dictTypeService) Delete(id int64) (err error) {
	err = global.GORM.Where("id = ?", id).Delete(&models.SysDictType{}).Error
	return
}

func (d dictTypeService) FindOne(id int64) (dictTypes *models.SysDictType, err error) {
	err = global.GORM.Where("id = ?", id).First(&dictTypes).Error
	return
}

func (d dictTypeService) Find() (dictTypes []*models.SysDictType, err error) {
	err = global.GORM.Find(&dictTypes).Error
	return
}
