package service

import (
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
)

type UserService struct{}

type IUserService interface {
	Create(user *models.SysUser)
	Find() []*models.SysUser
	FindOneByUserName(user *models.SysUser) error
}

func (u *UserService) Create(user *models.SysUser) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Find() (users []*models.SysUser) {
	global.GORM.Find(&users)
	return
}

func (u *UserService) FindOneByUserName(user *models.SysUser) error {
	if err := global.GORM.First(&user, "username").Error; err != nil {
		return err
	}
	return nil
}
