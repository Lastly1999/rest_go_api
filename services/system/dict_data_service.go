package service

import (
	request "resetgoapi.com/go-admin-api/common/request/system"
	"resetgoapi.com/go-admin-api/global"
	models "resetgoapi.com/go-admin-api/models/system"
	"resetgoapi.com/go-admin-api/pkg/db/scopes"
)

var DictDataService = dictDataService{}

type dictDataService struct{}

type IDictDataService interface {
	Page(request *request.DictDataListRequest) (*models.SysDictData, int64, error)
	Create(request *request.CreateDictDataRequest) error
	Update(request *request.UpdateDictDataRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysDictData, error)
	Find() ([]*models.SysDictData, error)
}

func (d dictDataService) Page(request *request.DictDataListRequest) (dictTypes *models.SysDictData, total int64, err error) {
	err = global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&dictTypes).Count(&total).Error
	return
}

func (d dictDataService) Create(request *request.CreateDictDataRequest) (err error) {
	err = global.GORM.Create(&models.SysDictData{
		DictLabel: request.DictLabel,
		DictValue: request.DictValue,
		DictSort:  request.DictSort,
		CreateBy:  request.CreateBy,
		Status:    request.Status,
		Remark:    request.Remark,
	}).Error
	return
}

func (d dictDataService) Update(request *request.UpdateDictDataRequest) (err error) {
	var dictData models.SysDictData
	dictData.Id = request.Id
	err = global.GORM.Model(&dictData).Updates(&models.SysDictData{
		DictLabel: request.DictLabel,
		DictValue: request.DictValue,
		DictSort:  request.DictSort,
		CreateBy:  request.CreateBy,
		Status:    request.Status,
		Remark:    request.Remark,
	}).Error
	return
}

func (d dictDataService) Delete(id int64) (err error) {
	err = global.GORM.Where("id = ?", id).Delete(&models.SysDictData{}).Error
	return
}

func (d dictDataService) FindOne(id int64) (dictTypes *models.SysDictData, err error) {
	err = global.GORM.Where("id = ?", id).First(&dictTypes).Error
	return
}

func (d dictDataService) Find() (dictTypes []*models.SysDictData, err error) {
	err = global.GORM.Find(&dictTypes).Error
	return
}
