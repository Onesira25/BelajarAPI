package model

import (
	"gorm.io/gorm"
)

type User struct {
	HP       string `json:"hp" form:"hp" validate:"required,min=10,max=13,number" gorm:"primaryKey"`
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Password string `json:"password" form:"password" validate:"required,alphanum"`
}

type ToDoList struct {
	UserHP      string  `json:"userhp" form:"userhp" gorm:"references:HP"`
	TaskName    string  `json:"taskname" form:"taskname" validate:"required"`
	DueDate     string  `json:"duedate" form:"duedate" validate:"required"`
	Description *string `json:"desc" form:"desc"`
}

type UserModel struct {
	Connection *gorm.DB
}

func (um *UserModel) Register(newData User) error {
	if err := um.Connection.Create(&newData).Error; err != nil {
		return err
	}

	return nil
}

func (um *UserModel) Login(hp string, password string) (User, error) {
	var result User
	if err := um.Connection.Where("hp = ? AND password = ?", hp, password).First(&result).Error; err != nil {
		return User{}, err
	}
	return result, nil
}

func (um *UserModel) AddTask(hp string, task ToDoList) error {
	newTask := ToDoList{
		UserHP:      task.UserHP,
		TaskName:    task.TaskName,
		DueDate:     task.DueDate,
		Description: task.Description,
	}

	if err := um.Connection.Create(&newTask).Error; err != nil {
		return err
	}
	return nil
}

func (um *UserModel) UpdateTask(hp string, updateTask ToDoList) (ToDoList, error) {
	var task ToDoList
	if err := um.Connection.Where("userhp = ?", updateTask.UserHP).First(&task).Error; err != nil {
		return task, err
	}

	task.UserHP = updateTask.UserHP

	if updateTask.TaskName != "" {
		task.TaskName = updateTask.TaskName
	}
	if updateTask.DueDate != "" {
		task.DueDate = updateTask.DueDate
	}
	if *updateTask.Description != "" {
		task.Description = updateTask.Description
	}

	if err := um.Connection.Save(&task).Error; err != nil {
		return ToDoList{}, err
	}

	return task, nil
}

func (um *UserModel) SeeAllTask() ([]ToDoList, error) {
	var allTask []ToDoList
	if err := um.Connection.Find(&allTask).Error; err != nil {
		return nil, err
	}

	return allTask, nil
}

// func (um *UserModel) CheckUser(id int) bool {
// 	var data User
// 	if err := um.Connection.Where("id = ?", id).First(&data).Error; err != nil {
// 		return false
// 	}
// 	return true
// }
