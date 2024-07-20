package service

import (
	"errors"
	request "resetgoapi.com/rest_go_api/common/request/system"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/utils/cypher"
)

type LoginService struct{}

type ILoginService interface {
	Sign(sign *request.SignRequest) (*request.SignResponse, error)
}

func (service *LoginService) Sign(sign *request.SignRequest) (*request.SignResponse, error) {
	userService := UserService{}
	user := &models.SysUser{
		Username: sign.Username,
	}
	if err := userService.FindOneByUserName(user); err != nil {
		return nil, err
	}
	if user == nil || !cypher.ComparePasswords(user.Password, sign.Password) {
		return nil, errors.New("账号或者密码错误，请重新登录！")
	}
	token := JwtService.GetToken(user)
	if token == "" {
		return nil, errors.New("token派发失败，请重试")
	}
	var resp request.SignResponse
	resp.AccessToken = token
	return &resp, nil
}
