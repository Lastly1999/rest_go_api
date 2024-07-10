package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resetgoapi.com/rest_go_api/pkg/config"
)

func Test(ctx *gin.Context) {
	ctx.String(http.StatusOK, config.GlobalConfig.App.Mode)
}
