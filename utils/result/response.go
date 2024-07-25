package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resetgoapi.com/go-admin-api/common/response"
)

type HttpResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type JsonResult struct {
	Context *gin.Context
}

func (jsonResult *JsonResult) HttpResultAuthError(msg string) {
	result := HttpResult{
		Code: HTTP_STATUS_AUTH_ERROR,
		Msg:  msg,
		Data: nil,
	}
	jsonResult.Context.JSON(http.StatusOK, result)
	jsonResult.Context.Abort()
}

func (jsonResult *JsonResult) HttpResultSuccess(data interface{}) {
	result := HttpResult{
		Code: HTTP_STATUS_SUCCESS,
		Msg:  "success",
		Data: data,
	}
	jsonResult.Context.JSON(http.StatusOK, result)
	jsonResult.Context.Abort()
}

func (jsonResult *JsonResult) HttpResultSuccessPage(data interface{}, total int64) {
	result := HttpResult{
		Code: HTTP_STATUS_SUCCESS,
		Msg:  "success",
		Data: &response.PageResponse{
			Total: total,
			List:  data,
		},
	}
	jsonResult.Context.JSON(http.StatusOK, result)
	jsonResult.Context.Abort()
}

func (jsonResult *JsonResult) HttpResultError(msg string) {
	result := HttpResult{
		Code: HTTP_STATUS_ERROR,
		Msg:  msg,
		Data: nil,
	}
	jsonResult.Context.JSON(http.StatusOK, result)
	jsonResult.Context.Abort()
}
