package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resetgoapi.com/rest_go_api/common/response"
)

type HttpResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func HttpResultAuthError(ctx *gin.Context, msg string) {
	result := HttpResult{
		Code: HTTP_STATUS_AUTH_ERROR,
		Msg:  msg,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, result)
	ctx.Abort()
}

func HttpResultSuccess(ctx *gin.Context, data interface{}) {
	result := HttpResult{
		Code: HTTP_STATUS_SUCCESS,
		Msg:  "success",
		Data: data,
	}
	ctx.JSON(http.StatusOK, result)
	ctx.Abort()
}

func HttpResultSuccessPage(ctx *gin.Context, data interface{}, total int64) {
	result := HttpResult{
		Code: HTTP_STATUS_SUCCESS,
		Msg:  "success",
		Data: &response.PageResponse{
			Total: total,
			List:  data,
		},
	}
	ctx.JSON(http.StatusOK, result)
	ctx.Abort()
}

func HttpResultError(ctx *gin.Context, msg string) {
	result := HttpResult{
		Code: HTTP_STATUS_ERROR,
		Msg:  msg,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, result)
	ctx.Abort()
}
