package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/rest_go_api/common/request/system"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/param"
	"resetgoapi.com/rest_go_api/utils/result"
)

type DictDataApi struct{}

// Create godoc
//
//	@Summary	创建字典数据
//	@Tags		数据字典数据
//	@Param		request	body	request.CreateDictDataRequest	true	"创建字典数据"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dictData/create [post]
func (api *DictDataApi) Create(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	createDictDataRequest := request.CreateDictDataRequest{}
	if err := ctx.ShouldBindJSON(createDictDataRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DictDataService.Create(&createDictDataRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Update godoc
//
//	@Summary	更新字典数据
//	@Tags		数据字典数据
//	@Param		request	body	request.updateDictDataRequest	true	"更新字典数据"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dictData/update [patch]
func (api *DictDataApi) Update(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	updateDictDataRequest := request.UpdateDictDataRequest{}
	if err := ctx.ShouldBindJSON(&updateDictDataRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DictDataService.Update(&updateDictDataRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Page godoc
//
//	@Summary	字典数据列表
//	@Tags		数据字典数据
//	@Accept		json
//	@Produce	json
//	@Param		userRequest	query		request.DictDataListRequest														false	"查询参数"
//	@Success	200			{object}	result.HttpResult{data=response.PageResponse{list=[]models.SysDictData}}	"desc"
//	@Router		/dictData/page [get]
func (api *DictDataApi) Page(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	dictDataListRequest := request.DictDataListRequest{}
	if err := ctx.ShouldBindJSON(&dictDataListRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	posts, total, err := service.DictDataService.Page(&dictDataListRequest)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccessPage(posts, total)
}

// Delete godoc
//
//	@Summary	删除字典数据
//	@Tags		数据字典数据
//	@Param		id	path		int					true	"字典数据id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dictData/delete/{id} [delete]
func (api *DictDataApi) Delete(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DictDataService.Delete(id); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Info godoc
//
//	@Summary	字典数据详情
//	@Tags		数据字典数据
//	@Param		id	path		int					true	"字典数据id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dictData/info/{id} [get]
func (api *DictDataApi) Info(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	user, err := service.DictDataService.FindOne(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(user)
}
