package main

import (
	"resetgoapi.com/rest_go_api/pkg/config"
	"resetgoapi.com/rest_go_api/pkg/db"
	"resetgoapi.com/rest_go_api/pkg/logger"
	"resetgoapi.com/rest_go_api/router"
)

func init() {
	config.Setup()
	db.Setup()
	logger.Setup()
}

func main() {
	router.InitRouter().Run()
}
