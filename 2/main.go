package main

import (
	"stockbit/2/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/get", controller.Get)

	e.Logger.Fatal(e.Start(":7000"))

}
