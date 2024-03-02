package main

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/user"
	"BelajarAPI/model"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	m := model.UserModel{Connection: db}
	c := user.UserController{Model: m}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, c)
	e.Logger.Fatal(e.Start(":8000"))
}
