package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	dsn := "root:@tcp(127.0.0.1:3306)/hotel?charset=utf8&parseTime=True&loc=Local"


	d, err := gorm.Open("mysql", dsn)

	if err != nil{
		log.Fatal("failed to connect to database:", err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}