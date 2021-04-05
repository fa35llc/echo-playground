package handler

import (
	"net/http"

	"github.com/bitclay/echo-playground/model"
	"github.com/labstack/echo/v4"
)

// curl -X POST http://localhost:80/create-user \
//  -H 'Content-Type: application/json' \
//  -d '{"name":"Joe","email":"joe@invalid-domain.com"}'
func CreateUser(c echo.Context) (err error) {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u.CreateUser(c)
	return c.JSON(http.StatusOK, u)
}
