package service

import (
	request "resetgoapi.com/rest_go_api/common/request/system"
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/pkg/db/scopes"
	"resetgoapi.com/rest_go_api/utils/cypher"
)

type UserService struct{}

type IUserService interface {
	Create(req *request.CreateUserRequest) error
	Update(req *request.UpdateUserRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysUser, error)
	FindPage(request *request.UserRequest) ([]*models.SysUser, int64)
	FindOneByUserName(user *models.SysUser) error
}

func (u *UserService) Delete(id int64) (err error) {
	err = global.GORM.Where("id = ?", id).Delete(&models.SysUser{}).Error
	return
}

func (u *UserService) FindOne(id int64) (user *models.SysUser, err error) {
	err = global.GORM.Where("id = ?", id).First(&user).Error
	return
}

func (u *UserService) Update(req *request.UpdateUserRequest) error {
	var user models.SysUser
	user.Id = req.Id
	if err := global.GORM.Model(&user).Updates(&models.SysUser{
		Username: req.Username,
		NickName: req.NickName,
		Email:    req.Email,
		Remark:   req.Remark,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
		Phone:    req.Phone,
		Status:   req.Status,
		DelFlag:  req.DelFlag,
		DeptId:   req.DeptId,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserService) Create(req *request.CreateUserRequest) error {
	password, _ := cypher.HashPassword("1234")
	if err := global.GORM.Create(&models.SysUser{
		Password: password,
		Username: req.Username,
		NickName: req.NickName,
		Email:    req.Email,
		Remark:   req.Remark,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
		Phone:    req.Phone,
		Status:   req.Status,
		DelFlag:  req.DelFlag,
		DeptId:   req.DeptId,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserService) FindPage(request *request.UserRequest) (users []*models.SysUser, total int64) {
	global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&users).Count(&total)
	return
}

func (u *UserService) FindOneByUserName(user *models.SysUser) (err error) {
	err = global.GORM.Where("username = ?", &user.Username).First(&user).Error
	return
}
