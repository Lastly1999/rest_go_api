package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GORM *gorm.DB

func Setup() {
	dsn := "root:Chen1027!@tcp(124.70.50.247:3306)/board?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	GORM = db
}
