package model

import (
	"fmt"

	"github.com/bitclay/echo-playground/db"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type (
	User struct {
		gorm.Model
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
	// Handler struct {
	// 	// Db map[string]*User
	// 	Db *User
	// }
)

func (u *User) CreateUser(c echo.Context) error {
	db := db.DbManager()
	// db.AutoMigrate(&User{})

	if err := db.Create(&u).Error; err != nil {
		fmt.Println("Error Gorm: Create user")
	}

	return nil
}
