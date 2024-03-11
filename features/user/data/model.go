package data

import (
	"BelajarAPI/features/todo/data"
)

type User struct {
	HP       string          `json:"hp" form:"hp" validate:"required,min=10,max=13,number" gorm:"type:varchar(13);primaryKey"`
	Name     string          `json:"name" form:"name" validate:"required,min=3"`
	Password string          `json:"-" validate:"required,alphanum"`
	Todos    []data.ToDoList `gorm:"foreignKey:UserHP;references:HP"`
}
