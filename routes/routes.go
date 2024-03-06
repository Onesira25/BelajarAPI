package routes

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/todo"
	"BelajarAPI/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, tc todo.TodoController) {
	userRoute(c, ctl)
	todoRoute(c, tc)
}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/register", ctl.Register())
	c.POST("/login", ctl.Login())
	c.PUT("/updateuser", ctl.UpdateUser(), withJWTConfig())
	c.GET("/allusers", ctl.GetAllUsers(), withJWTConfig())
	c.GET("/myprofile", ctl.MyProfile(), withJWTConfig())
}

func todoRoute(c *echo.Echo, tc todo.TodoController) {
	c.POST("/addtask", tc.AddTask(), withJWTConfig())
	c.PUT("/updatetask/:todoID", tc.UpdateTask(), withJWTConfig())
	c.GET("/alltask", tc.SeeAllMyTask(), withJWTConfig())

}

func withJWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	})
}
