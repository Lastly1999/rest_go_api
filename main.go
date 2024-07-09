package main

import (
	"resetgoapi.com/rest_go_api/pkg/db"
	"resetgoapi.com/rest_go_api/router"
)

func init() {
	db.Setup()
}

func main() {
	router.InitRouter().Run()
}
