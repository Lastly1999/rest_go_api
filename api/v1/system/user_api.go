package v1

import (
	"github.com/gin-gonic/gin"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/result"
)

type UserApi struct {
}

func (api *UserApi) Create(ctx *gin.Context) {

}

func (api *UserApi) Update(ctx *gin.Context) {

}

// Page godoc
// @Summary      用户列表
// @Description  get string by ID
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Router       /user/page [get]
func (api *UserApi) Page(ctx *gin.Context) {
	userService := service.UserService{}
	result.HttpResultSuccess(ctx, userService.Find())
}

func (api *UserApi) Delete(ctx *gin.Context) {

}
