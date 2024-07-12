package cmd

import (
	"fmt"
	"log"
	"net/http"
	_ "resetgoapi.com/rest_go_api/docs"
	"resetgoapi.com/rest_go_api/pkg/cache"
	"resetgoapi.com/rest_go_api/pkg/config"
	"resetgoapi.com/rest_go_api/pkg/db"
	"resetgoapi.com/rest_go_api/pkg/logger"
	"resetgoapi.com/rest_go_api/routers"
)

// Start godoc
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @BasePath  /v1
func Start() {
	logger.Setup()
	config.Setup()
	port := fmt.Sprintf(":%d", config.GlobalConfig.App.Port)
	// 启动服务器
	go func() {
		if err := routers.InitRouter().Run(port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 在后台执行初始化任务
	go func() {
		cache.Setup()
		db.Setup()
	}()

	// 主线程等待
	select {}
}
