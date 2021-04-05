package main

import (
	"github.com/bitclay/echo-playground/db"
	"github.com/bitclay/echo-playground/route"
	service "github.com/bitclay/echo-playground/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.Init()
	e := route.Init()
	e.HideBanner = true

	// Validator for non-main package
	e.Validator = &service.CustomValidator{Validator: validator.New()}

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Fix Allows Domain
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// e.Logger.Fatal(e.Start(":1323"))
	e.Start(":80")
}
