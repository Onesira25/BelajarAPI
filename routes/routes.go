package routes

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/register", ctl.Register())
	c.POST("/login", ctl.Login())
	c.POST("/addtask", ctl.AddTask(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/update", ctl.UpdateTask(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/alltask", ctl.SeeAllTask(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
