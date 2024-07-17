package service

import (
	request "resetgoapi.com/rest_go_api/common/request/system"
	models "resetgoapi.com/rest_go_api/models/system"
)

type LoginService struct{}

type ILoginService interface {
	Sign(sign *request.SignRequest) *request.SignResponse
}

func (service *LoginService) Sign(sign *request.SignRequest) (*request.SignResponse, error) {
	userSerivce := UserService{}
	user := models.SysUser{
		Username: sign.Username,
		Password: sign.Password,
	}
	if err := userSerivce.FindOneByUserName(&user); err != nil {
		return nil, err
	}
}
