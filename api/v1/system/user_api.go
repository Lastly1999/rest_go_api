package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/rest_go_api/common/request/system"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/param"
	"resetgoapi.com/rest_go_api/utils/result"
)

type UserApi struct {
}

// Create godoc
//
//	@Summary	创建用户
//	@Tags		用户管理
//	@Param		request	body	request.CreateUserRequest	true	"创建用户"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/user/create [post]
func (api *UserApi) Create(ctx *gin.Context) {
	createUserRequest := &request.CreateUserRequest{}
	jsonResult := result.JsonResult{Context: ctx}
	if err := ctx.ShouldBindJSON(createUserRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	userService := service.UserService{}
	if err := userService.Create(createUserRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Update godoc
//
//	@Summary	更新用户信息
//	@Tags		用户管理
//	@Accept		json
//	@Param		request	body	request.UpdateUserRequest	true	"更新用户"
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/user/update [patch]
func (api *UserApi) Update(ctx *gin.Context) {
	updateUserRequest := &request.UpdateUserRequest{}
	jsonResult := result.JsonResult{Context: ctx}
	if err := ctx.ShouldBindJSON(updateUserRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	userService := service.UserService{}
	if err := userService.Update(updateUserRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Page godoc
//
//	@Summary	用户列表
//	@Tags		用户管理
//	@Accept		json
//	@Produce	json
//	@Param		userRequest	query		request.UserRequest														false	"查询参数"
//	@Success	200			{object}	result.HttpResult{data=response.PageResponse{list=[]models.SysUser}}	"desc"
//	@Router		/user/page [get]
func (api *UserApi) Page(ctx *gin.Context) {
	userReq := &request.UserRequest{}
	jsonResult := result.JsonResult{Context: ctx}
	if err := ctx.ShouldBindQuery(userReq); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	userService := service.UserService{}
	list, total := userService.FindPage(userReq)
	jsonResult.HttpResultSuccessPage(list, total)
}

// Delete godoc
//
//	@Summary	删除用户
//	@Tags		用户管理
//	@Param		id	path		int					true	"用户id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/user/delete/{id} [delete]
func (api *UserApi) Delete(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	userService := service.UserService{}
	err = userService.Delete(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Info godoc
//
//	@Summary	用户详情
//	@Tags		用户管理
//	@Param		id	path		int										true	"用户id"
//	@Success	200	{object}	result.HttpResult{data=models.SysUser}	"desc"
//	@Router		/user/info/{id} [get]
func (api *UserApi) Info(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	userService := service.UserService{}
	user, err := userService.FindOne(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(user)
}
