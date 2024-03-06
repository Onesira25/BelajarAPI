package main

import (
	"BelajarAPI/config"
	tControl "BelajarAPI/controller/todo"
	uControl "BelajarAPI/controller/user"
	"BelajarAPI/model/todo"
	"BelajarAPI/model/user"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	m := user.UserModel{Connection: db}
	c := uControl.UserController{Model: m}
	tm := todo.ToDoModel{Connection: db}
	tc := tControl.TodoController{Model: tm}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, c, tc)
	e.Logger.Fatal(e.Start(":8000"))
}
