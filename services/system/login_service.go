package service

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/rest_go_api/common/request/system"
)

type LoginService struct{}

type ILoginService interface {
	Sign(sign *request.SignRequest) *gin.ResponseWriter
}

func (service *LoginService) Sign(sign *request.SignRequest) *gin.ResponseWriter {
	//TODO implement me
	panic("implement me")
}
