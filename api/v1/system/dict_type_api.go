package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/rest_go_api/common/request/system"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/param"
	"resetgoapi.com/rest_go_api/utils/result"
)

type DictTypeApi struct{}

// Create godoc
//
//	@Summary	创建字典类型
//	@Tags		数据字典类型
//	@Param		request	body	request.CreateDictTypeRequest	true	"创建字典类型"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dictType/create [post]
func (api *DictTypeApi) Create(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	createDictTypeRequest := request.CreateDictTypeRequest{}
	if err := ctx.ShouldBindJSON(createDictTypeRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DictTypeService.Create(&createDictTypeRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Update godoc
//
//	@Summary	更新字典类型
//	@Tags		数据字典类型
//	@Param		request	body	request.UpdateDictTypeRequest	true	"更新字典类型"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dictType/update [patch]
func (api *DictTypeApi) Update(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	updateDictTypeRequest := request.UpdateDictTypeRequest{}
	if err := ctx.ShouldBindJSON(&updateDictTypeRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DictTypeService.Update(&updateDictTypeRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Page godoc
//
//	@Summary	字典类型列表
//	@Tags		数据字典类型
//	@Accept		json
//	@Produce	json
//	@Param		userRequest	query		request.DictTypeListRequest														false	"查询参数"
//	@Success	200			{object}	result.HttpResult{data=response.PageResponse{list=[]models.SysDictType}}	"desc"
//	@Router		/dictType/page [get]
func (api *DictTypeApi) Page(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	dictTypeListRequest := request.DictTypeListRequest{}
	if err := ctx.ShouldBindJSON(&dictTypeListRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	posts, total, err := service.DictTypeService.Page(&dictTypeListRequest)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccessPage(posts, total)
}

// Delete godoc
//
//	@Summary	删除字典类型
//	@Tags		数据字典类型
//	@Param		id	path		int					true	"字典类型id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dictType/delete/{id} [delete]
func (api *DictTypeApi) Delete(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DictTypeService.Delete(id); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Info godoc
//
//	@Summary	字典类型详情
//	@Tags		数据字典类型
//	@Param		id	path		int					true	"字典类型id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dictType/info/{id} [get]
func (api *DictTypeApi) Info(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	user, err := service.DictTypeService.FindOne(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(user)
}
