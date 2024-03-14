package todo

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ToDoController interface {
	AddTask() echo.HandlerFunc
	UpdateTask() echo.HandlerFunc
	SeeAllMyTask() echo.HandlerFunc
}

type ToDoModel interface {
	AddTask(hp string, task ToDo) (ToDo, error)
	UpdateTask(hp string, todoID uint, updateTask ToDo) (ToDo, error)
	SeeAllMyTask(hp string) ([]ToDo, error)
}

type TodoServices interface {
	AddTask(token *jwt.Token, task ToDo) (ToDo, error)
	UpdateTask(token *jwt.Token, todoID string, updateTask ToDo) (ToDo, error)
	SeeAllMyTask(token *jwt.Token) ([]ToDo, error)
}

type ToDo struct {
	gorm.Model
	UserHP      string
	TaskName    string
	DueDate     string
	Description string
}

type AddToDo struct {
	UserHP      string `json:"userhp" form:"userhp" validate:"required,min=10,max=13" gorm:"type:varchar(13)"`
	TaskName    string `json:"taskname" form:"taskname" validate:"required"`
	DueDate     string `json:"duedate" form:"duedate" validate:"required"`
	Description string `json:"desc" form:"desc"`
}

type UpdateTodo struct {
	TaskName    string `json:"taskname" form:"taskname" validate:"required"`
	DueDate     string `json:"duedate" form:"duedate" validate:"required"`
	Description string `json:"desc" form:"desc"`
}
