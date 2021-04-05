package model

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type (
	Product struct {
		gorm.Model
		Name     string `json:"name" validate:"required"`
		Category string `json:"category" validate:"required"`
		Price    string `json:"price" validate:"required"`
	}
)

func (u *Product) GetUser(c echo.Context) (Product, error) {
	// name := c.Param("name")
	// email := c.Param("email")
	var product Product
	if err := c.Bind(&product); err != nil {
		//error handling
		return product, err
	}
	return product, nil
}
