package main

import (
	"fmt"
	"log"
	"net/http"
	"resetgoapi.com/rest_go_api/pkg/config"
	"resetgoapi.com/rest_go_api/pkg/db"
	"resetgoapi.com/rest_go_api/pkg/logger"
	"resetgoapi.com/rest_go_api/router"
)

func main() {
	logger.Setup()
	config.Setup()
	port := fmt.Sprintf(":%d", config.GlobalConfig.App.Port)
	// 启动服务器
	go func() {
		if err := router.InitRouter().Run(port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 在后台执行初始化任务
	go func() {
		db.Setup()
	}()

	// 主线程等待
	select {}
}
