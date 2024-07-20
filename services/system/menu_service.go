package service

import (
	request "resetgoapi.com/rest_go_api/common/request/system"
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/pkg/db/scopes"
)

var MenuService = menuService{}

type menuService struct{}

type IMenuService interface {
	Page(request *request.MenuListRequest) (*models.SysMenu, int64, error)
	Create(request *request.CreateMenuRequest) error
	Update(request *request.UpdateMenuRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysMenu, error)
	Find() ([]*models.SysMenu, error)
}

func (m menuService) Page(request *request.MenuListRequest) (menus *models.SysMenu, total int64, err error) {
	err = global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&menus).Count(&total).Error
	return
}

func (m menuService) Create(request *request.CreateMenuRequest) (err error) {
	err = global.GORM.Create(&models.SysMenu{
		MenuName:   request.MenuName,
		MenuIcon:   request.MenuIcon,
		ParentId:   request.ParentId,
		ParentName: request.ParentName,
		MenuSort:   request.MenuSort,
		MenuType:   request.MenuType,
		Visible:    request.Visible,
		Perms:      request.Perms,
		Remark:     request.Remark,
		IsCache:    request.IsCache,
		IsFrame:    request.IsFrame,
		Component:  request.Component,
		Router:     request.Router,
		Status:     request.Status,
	}).Error
	return
}

func (m menuService) Update(request *request.UpdateMenuRequest) (err error) {
	var menu models.SysMenu
	menu.Id = request.Id
	err = global.GORM.Model(&menu).Updates(&models.SysMenu{
		MenuName:   request.MenuName,
		MenuIcon:   request.MenuIcon,
		ParentId:   request.ParentId,
		ParentName: request.ParentName,
		MenuSort:   request.MenuSort,
		MenuType:   request.MenuType,
		Visible:    request.Visible,
		Perms:      request.Perms,
		Remark:     request.Remark,
		IsCache:    request.IsCache,
		IsFrame:    request.IsFrame,
		Component:  request.Component,
		Router:     request.Router,
		Status:     request.Status,
	}).Error
	return
}

func (m menuService) Delete(id int64) (err error) {
	err = global.GORM.Where("id = ?", id).Delete(&models.SysMenu{}).Error
	return
}

func (m menuService) FindOne(id int64) (menu *models.SysMenu, err error) {
	err = global.GORM.Where("id = ?", id).First(&menu).Error
	return
}

func (m menuService) Find() (menus []*models.SysMenu, err error) {
	err = global.GORM.Find(&menus).Error
	return
}
