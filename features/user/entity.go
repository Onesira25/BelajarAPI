package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	MyProfile() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
}
type UserModel interface {
	Register(newData User) error
	Login(hp string) (User, error)
	MyProfile(hp string) (User, error)
	UpdateUser(hp string, updateData User) error
	GetAllUsers(hp string) ([]User, error)
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	MyProfile(token *jwt.Token) (User, error)
	UpdateUser(token *jwt.Token, updateData User) error
	GetAllUsers(token *jwt.Token) ([]User, error)
}

type User struct {
	HP       string
	Name     string
	Password string
}

type Login struct {
	HP       string `json:"hp" form:"hp" validate:"required,min=10,max=13,number" gorm:"type:varchar(13);primaryKey"`
	Password string `json:"-" validate:"required,alphanum"`
}

type Register struct {
	HP       string `json:"hp" form:"hp" validate:"required,min=10,max=13,number" gorm:"type:varchar(13);primaryKey"`
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Password string `json:"-" validate:"required,alphanum"`
}

type UpdateUser struct {
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Password string `json:"-" form:"password" validate:"required,alphanum"`
}
