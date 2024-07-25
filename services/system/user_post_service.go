package service

import (
	"resetgoapi.com/go-admin-api/global"
	models "resetgoapi.com/go-admin-api/models/system"
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
