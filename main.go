package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thomaszub/go-url-shortener/api"
)

func main() {
	var controller api.ServerInterface = NewController()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	api.RegisterHandlersWithBaseURL(e, controller, "/api")
	e.Logger.Fatal(e.Start(":8080"))
}
