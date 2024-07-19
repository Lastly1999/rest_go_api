package service

import (
	"errors"
	request "resetgoapi.com/rest_go_api/common/request/system"
	models "resetgoapi.com/rest_go_api/models/system"
)

type LoginService struct{}

type ILoginService interface {
	Sign(sign *request.SignRequest) (*request.SignResponse, error)
}

func (service *LoginService) Sign(sign *request.SignRequest) (*request.SignResponse, error) {
	userService := UserService{}
	user := &models.SysUser{
		Username: sign.Username,
		Password: sign.Password,
	}
	if err := userService.FindOneByUserName(user); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("登录失败，请重试！")
	}
	jwtService := JwtService{}
	token := jwtService.GetToken(user)
	if token == "" {
		return nil, errors.New("token派发失败，请重试")
	}
	var resp request.SignResponse
	resp.AccessToken = token
	return &resp, nil
}
