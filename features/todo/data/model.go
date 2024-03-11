package data

import "gorm.io/gorm"

type ToDoList struct {
	gorm.Model
	UserHP      string `json:"userhp" form:"userhp" validate:"required,min=10,max=13" gorm:"type:varchar(13)"`
	TaskName    string `json:"taskname" form:"taskname" validate:"required"`
	DueDate     string `json:"duedate" form:"duedate" validate:"required"`
	Description string `json:"desc" form:"desc"`
}
