package main

import (
	"task/cmd/api/handlers/routes"
	"task/internal/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	configs.DatabaseInit()
	db := configs.DB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}
