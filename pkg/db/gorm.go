package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"resetgoapi.com/rest_go_api/pkg/config"
)

var GORM *gorm.DB

func Setup() {
	mysqlConfig := config.Config.Database.MysqlConfig
	dsn := fmt.Sprintf("%s:%s!@tcp(%s:%s)/board?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	GORM = db
}
