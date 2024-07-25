package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/go-admin-api/common/request/system"
	service "resetgoapi.com/go-admin-api/services/system"
	"resetgoapi.com/go-admin-api/utils/result"
)

type LoginApi struct {
}

// Login godoc
//
//	@Summary	用户登录
//	@Tags		授权
//	@Param		request	body		request.SignRequest								true	"用户登录"
//	@success	200		{object}	result.HttpResult{data=response.SignResponse}	"desc"
//	@Router		/login/sign [post]
func (api *LoginApi) Login(ctx *gin.Context) {
	signRequest := request.SignRequest{}
	jsonResult := result.JsonResult{Context: ctx}
	if err := ctx.ShouldBindJSON(&signRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	resp, err := service.LoginService.Sign(&signRequest)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(resp)
}
