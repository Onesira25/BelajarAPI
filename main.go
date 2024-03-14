package main

import (
	"BelajarAPI/config"
	td "BelajarAPI/features/todo/data"
	th "BelajarAPI/features/todo/handler"
	ts "BelajarAPI/features/todo/services"
	"BelajarAPI/features/user/data"
	"BelajarAPI/features/user/handler"
	"BelajarAPI/features/user/services"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	todoData := td.New(db)
	todoService := ts.NewTodoService(todoData)
	todoHandler := th.NewHandler(todoService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, userHandler, todoHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
