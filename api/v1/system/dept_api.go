package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/rest_go_api/common/request/system"
	models "resetgoapi.com/rest_go_api/models/system"
	service "resetgoapi.com/rest_go_api/services/system"
	"resetgoapi.com/rest_go_api/utils/param"
	"resetgoapi.com/rest_go_api/utils/result"
)

type DeptApi struct{}

// Page godoc
//
//	@Summary	部门列表分页
//	@Tags		部门管理
//	@Accept		json
//	@Produce	json
//	@Param		DeptListRequest	query		request.DeptListRequest														false	"查询参数"
//	@Success	200			{object}	result.HttpResult{data=response.PageResponse{list=[]models.SysDept}}	"desc"
//	@Router		/dept/page [get]
func (api *DeptApi) Page(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	deptListRequest := request.DeptListRequest{}
	if err := ctx.ShouldBind(&deptListRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	depts, err := service.DeptService.Page(&deptListRequest)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	treeDepts := models.BuildDeptTree(depts)
	jsonResult.HttpResultSuccess(treeDepts)
}

// List godoc
//
//	@Summary	部门列表
//	@Tags		部门管理
//	@Accept		json
//	@Produce	json
//	@Param		DeptListRequest	query		request.DeptListRequest														false	"查询参数"
//	@Success	200			{object}	result.HttpResult{data=response.PageResponse{list=[]models.SysDept}}	"desc"
//	@Router		/dept/page [get]
func (api *DeptApi) List(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	list, err := service.DeptService.List()
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(list)
}

// Create godoc
//
//	@Summary	创建部门
//	@Tags		部门管理
//	@Param		request	body	request.CreateDeptRequest	true	"创建部门"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dept/create [post]
func (api *DeptApi) Create(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	createDeptRequest := request.CreateDeptRequest{}
	if err := ctx.ShouldBindJSON(&createDeptRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DeptService.Create(&createDeptRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Update godoc
//
//	@Summary	更新部门
//	@Tags		部门管理
//	@Param		request	body	request.UpdateDeptRequest	true	"更新部门"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dept/update [patch]
func (api *DeptApi) Update(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	updateDeptRequest := request.UpdateDeptRequest{}
	if err := ctx.ShouldBindJSON(&updateDeptRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DeptService.Update(&updateDeptRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Delete godoc
//
//	@Summary	删除部门
//	@Tags		部门管理
//	@Param		id	path		int					true	"部门id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/dept/delete/{id} [delete]
func (api *DeptApi) Delete(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.DeptService.Delete(id); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Info godoc
//
//	@Summary	部门详情
//	@Tags		部门管理
//	@Param		id	path		int					true	"部门id"
//	@Success	200	{object}	result.HttpResult{data=models.SysDept}	"desc"
//	@Router		/dept/info/{id} [get]
func (api *DeptApi) Info(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	id, err := param.GetParamInt64(ctx, "id", 0)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	dept, err := service.DeptService.FindOne(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(dept)
}
