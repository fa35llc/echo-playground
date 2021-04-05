package db

import (
	"fmt"

	"github.com/bitclay/echo-playground/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	configuration := config.GetConfig()
	fmt.Println(configuration)
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connect_string)
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	// db.AutoMigrate(&model.User{})
	return db
}

func DbManager() *gorm.DB {
	return db
}
