package data

import (
	"BelajarAPI/features/todo"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) todo.ToDoModel {
	return &model{
		connection: db,
	}
}

func (tm *model) AddTask(hp string, task todo.ToDo) (todo.ToDo, error) {
	if err := tm.connection.Create(&task).Error; err != nil {
		return todo.ToDo{}, err
	}
	return task, nil
}

func (tm *model) UpdateTask(hp string, todoid uint, updateTask todo.ToDo) (todo.ToDo, error) {
	qry := tm.connection.Where("user_hp = ? AND id = ?", hp, todoid).Updates(&updateTask)
	if err := qry.Error; err != nil {
		return todo.ToDo{}, err
	}

	if qry.RowsAffected < 1 {
		return todo.ToDo{}, errors.New("no data affected")
	}

	return updateTask, nil
}

func (tm *model) SeeAllMyTask(userhp string) ([]todo.ToDo, error) {
	var allTask []todo.ToDo
	if err := tm.connection.Where("user_hp = ?", userhp).Find(&allTask).Error; err != nil {
		return nil, err
	}

	return allTask, nil
}
