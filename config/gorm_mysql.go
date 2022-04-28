package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	DB_URL string = "root:root@tcp(127.0.0.1:3306)/go-test?charset=utf8mb4&parseTime=True&loc=Local"
)

func Connect() {
	d, err := gorm.Open(mysql.Open(DB_URL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Mysql connected!!!")
	db = d
}

func getDb() *gorm.DB {
	return db
}
