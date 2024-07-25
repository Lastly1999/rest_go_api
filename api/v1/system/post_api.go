package v1

import (
	"github.com/gin-gonic/gin"
	request "resetgoapi.com/go-admin-api/common/request/system"
	service "resetgoapi.com/go-admin-api/services/system"
	"resetgoapi.com/go-admin-api/utils/param"
	"resetgoapi.com/go-admin-api/utils/result"
)

type PostApi struct{}

// Create godoc
//
//	@Summary	创建岗位
//	@Tags		岗位管理
//	@Param		request	body	request.CreatePostRequest	true	"创建岗位"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/post/create [post]
func (api *PostApi) Create(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	createPostRequest := request.CreatePostRequest{}
	if err := ctx.ShouldBindJSON(createPostRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.PostService.Create(&createPostRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Update godoc
//
//	@Summary	更新岗位
//	@Tags		岗位管理
//	@Param		request	body	request.UpdateRoleRequest	true	"更新岗位"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/post/update [patch]
func (api *PostApi) Update(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	updatePostRequest := request.UpdatePostRequest{}
	if err := ctx.ShouldBindJSON(&updatePostRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.PostService.Update(&updatePostRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Page godoc
//
//	@Summary	岗位列表
//	@Tags		岗位管理
//	@Accept		json
//	@Produce	json
//	@Param		userRequest	query		request.PostListRequest														false	"查询参数"
//	@Success	200			{object}	result.HttpResult{data=response.PageResponse{list=[]models.SysPost}}	"desc"
//	@Router		/post/page [get]
func (api *PostApi) Page(ctx *gin.Context) {
	jsonResult := result.JsonResult{Context: ctx}
	postListRequest := request.PostListRequest{}
	if err := ctx.ShouldBindJSON(&postListRequest); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	posts, total, err := service.PostService.Page(&postListRequest)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccessPage(posts, total)
}

// Delete godoc
//
//	@Summary	删除岗位
//	@Tags		岗位管理
//	@Param		id	path		int					true	"岗位id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/post/delete/{id} [delete]
func (api *PostApi) Delete(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	if err := service.PostService.Delete(id); err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(nil)
}

// Info godoc
//
//	@Summary	岗位详情
//	@Tags		岗位管理
//	@Param		id	path		int					true	"岗位id"
//	@Success	200	{object}	result.HttpResult	"desc"
//	@Router		/post/info/{id} [get]
func (api *PostApi) Info(ctx *gin.Context) {
	id, err := param.GetParamInt64(ctx, "id", 0)
	jsonResult := result.JsonResult{Context: ctx}
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	user, err := service.PostService.FindOne(id)
	if err != nil {
		jsonResult.HttpResultError(err.Error())
		return
	}
	jsonResult.HttpResultSuccess(user)
}
