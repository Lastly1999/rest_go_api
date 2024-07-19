package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/rest_go_api/common/request/system"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/result"
)

type LoginApi struct {
}

// Login godoc
//
//	@Summary	用户登录
//	@Tags		授权
//	@Param		request	body		request.SignRequest								true	"用户登录"
//	@success	200		{object}	result.HttpResult{data=request.SignResponse}	"desc"
//	@Router		/login/sign [post]
func (api *LoginApi) Login(ctx *gin.Context) {
	signRequest := request.SignRequest{}
	jsonResult := result.JsonResult{Context: ctx}
	if err := ctx.ShouldBindJSON(&signRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	loginService := service.LoginService{}
	resp, err := loginService.Sign(&signRequest)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(resp)
}
