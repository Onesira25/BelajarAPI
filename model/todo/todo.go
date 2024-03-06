package todo

import (
	"errors"

	"gorm.io/gorm"
)

type ToDoList struct {
	gorm.Model
	UserHP      string `json:"userhp" form:"userhp" validate:"required,min=10,max=13" gorm:"foreignKey;references:HP"`
	TaskName    string `json:"taskname" form:"taskname" validate:"required"`
	DueDate     string `json:"duedate" form:"duedate" validate:"required"`
	Description string `json:"desc" form:"desc"`
}

type ToDoModel struct {
	Connection *gorm.DB
}

func (tm *ToDoModel) AddTask(task ToDoList) (ToDoList, error) {
	if err := tm.Connection.Create(&task).Error; err != nil {
		return ToDoList{}, err
	}
	return task, nil
}

func (tm *ToDoModel) UpdateTask(hp string, todoid uint, updateTask ToDoList) (ToDoList, error) {
	qry := tm.Connection.Where("user_hp = ? AND id = ?", hp, todoid).Updates(&updateTask)
	if err := qry.Error; err != nil {
		return ToDoList{}, err
	}

	if qry.RowsAffected < 1 {
		return ToDoList{}, errors.New("no data affected")
	}

	return updateTask, nil
}

func (tm *ToDoModel) SeeAllTask(userhp string) ([]ToDoList, error) {
	var allTask []ToDoList
	if err := tm.Connection.Where("user_hp = ?", userhp).Find(&allTask).Error; err != nil {
		return nil, err
	}

	return allTask, nil
}
