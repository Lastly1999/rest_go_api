package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/rest_go_api/common/request/system"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/param"
	"resetgoapi.com/rest_go_api/utils/result"
)

type MenuApi struct{}

// Create godoc
//
//	@Summary	创建菜单
//	@Tags		菜单管理
//	@Param		request	body	request.CreateMenuRequest	true	"创建菜单"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/menu/create [post]
func (api *MenuApi) Create(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	createMenuRequest := request.CreateMenuRequest{}
	if err := ctx.ShouldBindJSON(&createMenuRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.MenuService.Create(&createMenuRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Update godoc
//
//	@Summary	创建菜单
//	@Tags		菜单管理
//	@Param		request	body	request.UpdateMenuRequest	true	"创建菜单"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/menu/update [patch]
func (api *MenuApi) Update(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	updateMenuRequest := request.UpdateMenuRequest{}
	if err := ctx.ShouldBindJSON(&updateMenuRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.MenuService.Update(&updateMenuRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Info godoc
//
//	@Summary	菜单详情
//	@Tags		菜单管理
//	@Param		id	path		int					true	"菜单id"
//	@Success	200	{object}	result.HttpResult{data=models.SysMenu}	"desc"
//	@Router		/menu/info/{id} [get]
func (api *MenuApi) Info(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	id, err := param.GetParamInt64(ctx, "id", 0)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	user, err := service.MenuService.FindOne(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(user)
}

// Page godoc
//
//	@Summary	菜单列表
//	@Tags		菜单管理
//	@Accept		json
//	@Produce	json
//	@Param		userRequest	query		request.UserRequest														false	"查询参数"
//	@Success	200			{object}	result.HttpResult{data=response.PageResponse{list=[]models.SysMenu}}	"desc"
//	@Router		/menu/page [get]
func (api *MenuApi) Page(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	menuListRequest := request.MenuListRequest{}
	if err := ctx.ShouldBindQuery(&menuListRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	menus, total, err := service.MenuService.Page(&menuListRequest)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccessPage(menus, total)
}

// Delete godoc
//
//	@Summary	删除菜单
//	@Tags		菜单管理
//	@Param		id	path		int					true	"菜单id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/menu/delete/{id} [delete]
func (api *MenuApi) Delete(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	id, err := param.GetParamInt64(ctx, "id", 0)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	menu, err := service.MenuService.FindOne(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(menu)
}
