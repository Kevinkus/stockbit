package controller

import (
	"net/http"
	"stockbit/2/service"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	response, err := service.Get(c.QueryParam("pagination"), c.QueryParam("searchword"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
