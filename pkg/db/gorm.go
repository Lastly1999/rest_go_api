package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"resetgoapi.com/rest_go_api/global"
	"resetgoapi.com/rest_go_api/pkg/config"
)

func Setup() {
	mysqlConfig := config.GlobalConfig.Database.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/board?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Logger.Error("Failed to connect to the database. Please check your configuration and database server.")
	}
	global.GORM = db
	global.Logger.Info(`Successfully connected to the database. The database connection was established correctly, and the system is now ready to perform database operations.`)
}
