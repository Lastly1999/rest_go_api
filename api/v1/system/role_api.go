package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/rest_go_api/common/request/system"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/param"
	"resetgoapi.com/rest_go_api/utils/result"
)

type RoleApi struct{}

// Create godoc
//
//	@Summary	创建角色
//	@Tags		角色管理
//	@Param		request	body	request.CreateRoleRequest	true	"创建角色"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/role/create [post]
func (api *RoleApi) Create(ctx *gin.Context) {
	roleRequest := &request.CreateRoleRequest{}
	jsonResult := result.JsonResult{Context: ctx}
	if err := ctx.ShouldBindJSON(&roleRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.RoleService.Create(roleRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Update godoc
//
//	@Summary	更新角色
//	@Tags		角色管理
//	@Param		request	body	request.UpdateRoleRequest	true	"更新角色"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/role/create [patch]
func (api *RoleApi) Update(ctx *gin.Context) {
	roleRequest := request.UpdateRoleRequest{}
	jsonResult := result.JsonResult{Context: ctx}
	if err := ctx.ShouldBindJSON(&roleRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.RoleService.Update(&roleRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Delete godoc
//
//	@Summary	删除角色
//	@Tags		角色管理
//	@Param		id	path		int					true	"角色id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/role/delete/{id} [delete]
func (api *RoleApi) Delete(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.RoleService.Delete(id); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Info godoc
//
//	@Summary	角色详情
//	@Tags		角色管理
//	@Param		id	path		int					true	"角色id"
//	@Success	200	{object}	result.HttpResult{data=models.SysRole}	"desc"
//	@Router		/role/info/{id} [get]
func (api *RoleApi) Info(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	role, err := service.RoleService.FindOne(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(role)
}

// Page godoc
//
//	@Summary	角色列表
//	@Tags		角色管理
//	@Accept		json
//	@Produce	json
//	@Param		userRequest	query		request.UserRequest														false	"查询参数"
//	@Success	200			{object}	result.HttpResult{data=response.PageResponse{list=[]models.SysRole}}	"desc"
//	@Router		/role/page [get]
func (api *RoleApi) Page(ctx *gin.Context) {
	roleListRequest := &request.RoleListRequest{}
	jsonResult := result.JsonResult{Context: ctx}
	if err := ctx.ShouldBind(roleListRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	roles, total, err := service.RoleService.Page(roleListRequest)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccessPage(roles, total)
}
