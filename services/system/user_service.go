package service

import (
	request "resetgoapi.com/go-admin-api/common/request/system"
	response "resetgoapi.com/go-admin-api/common/response/system"
	"resetgoapi.com/go-admin-api/global"
	models "resetgoapi.com/go-admin-api/models/system"
	"resetgoapi.com/go-admin-api/pkg/db/scopes"
	"resetgoapi.com/go-admin-api/utils/cypher"
)

var UserService = userService{}

type userService struct{}

type IUserService interface {
	Create(req *request.CreateUserRequest) error
	Update(req *request.UpdateUserRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysUser, error)
	FindPage(request *request.UserRequest) ([]*models.SysUser, int64)
	FindOneByUserName(user *models.SysUser) error
}

func (u *userService) Delete(id int64) (err error) {
	err = global.GORM.Where("id = ?", id).Delete(&models.SysUser{}).Error
	return
}

func (u *userService) FindOne(id int64) (user *models.SysUser, err error) {
	err = global.GORM.Where("id = ?", id).First(&user).Error
	return
}

func (u *userService) Update(req *request.UpdateUserRequest) error {
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

func (u *userService) Create(req *request.CreateUserRequest) error {
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

func (u *userService) FindPage(request *request.UserRequest) ([]*response.SysUserPageListResponse, int64) {
	var users []*models.SysUser
	var total int64
	global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&users).Count(&total)
	var pageList []*response.SysUserPageListResponse
	for _, user := range users {
		ids, _ := UserRoleService.FindRoleIdsByUserId(user.Id)
		userRoles, _ := RoleService.FindRolesInIds(ids)
		var roleNames []string
		for _, role := range userRoles {
			roleNames = append(roleNames, role.RoleName)
		}
		pageList = append(pageList, &response.SysUserPageListResponse{
			Role:    roleNames,
			SysUser: *user,
		})
	}
	return pageList, total
}

func (u *userService) FindOneByUserName(user *models.SysUser) (err error) {
	err = global.GORM.Where("username = ?", &user.Username).First(&user).Error
	return
}
