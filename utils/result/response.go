package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func HttpResultSuccess(ctx *gin.Context, data interface{}) {
	result := HttpResult{
		Code: HTTP_STATUS_SUCCESS,
		Msg:  "success",
		Data: data,
	}
	ctx.JSON(http.StatusOK, result)
}

func HttpResultError(ctx *gin.Context, msg string) {
	result := HttpResult{
		Code: HTTP_STATUS_ERROR,
		Msg:  msg,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, result)
}
