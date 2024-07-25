package service

import (
	"errors"
	request "resetgoapi.com/go-admin-api/common/request/system"
	response "resetgoapi.com/go-admin-api/common/response/system"
	models "resetgoapi.com/go-admin-api/models/system"
	"resetgoapi.com/go-admin-api/utils/cypher"
)

var LoginService = loginService{}

type loginService struct{}

type ILoginService interface {
	Sign(sign *request.SignRequest) (*response.SignResponse, error)
	GetUserInfo(userId int64) (*response.GetUserInfoResponse, error)
}

func (service *loginService) Sign(sign *request.SignRequest) (*response.SignResponse, error) {
	user := &models.SysUser{
		Username: sign.Username,
	}
	if err := UserService.FindOneByUserName(user); err != nil {
		return nil, errors.New("账号或者密码错误，请重新登录！")
	}
	if user == nil || !cypher.ComparePasswords(user.Password, sign.Password) {
		return nil, errors.New("账号或者密码错误，请重新登录！")
	}
	token := JwtService.GetToken(user)
	if token == "" {
		return nil, errors.New("token派发失败，请重试")
	}
	var resp response.SignResponse
	resp.AccessToken = token
	return &resp, nil
}

func (service *loginService) GetUserInfo(userId int64) (*response.GetUserInfoResponse, error) {
	userInfo, err := UserService.FindOne(userId)
	if err != nil {
		return nil, err
	}
	return &response.GetUserInfoResponse{
		RealName: userInfo.NickName,
		Username: userInfo.Username,
		UserId:   userInfo.Id,
		Avatar:   userInfo.Avatar,
		Desc:     userInfo.Remark,
	}, nil
}
