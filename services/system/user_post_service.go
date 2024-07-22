package service

import (
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
)

var UserPostService = userPostService{}

type userPostService struct{}

type IUserPostService interface {
	Create(request *models.SysUserPost) error
}

func (u userPostService) Create(userPost *models.SysUserPost) (err error) {
	err = global.GORM.Create(&userPost).Error
	return
}
