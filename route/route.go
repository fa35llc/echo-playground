package route

import (
	"github.com/bitclay/echo-playground/handler"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.POST("/create-user", handler.CreateUser)
	return e
}
