package services

import (
	"BelajarAPI/features/todo"
	"BelajarAPI/helper"
	"BelajarAPI/middlewares"
	"errors"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type services struct {
	m todo.ToDoModel
	v *validator.Validate
}

func NewTodoService(model todo.ToDoModel) todo.TodoServices {
	return &services{
		m: model,
		v: validator.New(),
	}
}

func (s *services) AddTask(token *jwt.Token, newTodo todo.ToDo) (todo.ToDo, error) {
	decodeHP := middlewares.DecodeToken(token)
	if decodeHP == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return todo.ToDo{}, errors.New("data tidak valid")
	}

	var validateTask todo.AddToDo
	validateTask.UserHP = decodeHP
	validateTask.TaskName = newTodo.TaskName
	validateTask.DueDate = newTodo.DueDate
	validateTask.Description = newTodo.Description
	err := s.v.Struct(&validateTask)
	if err != nil {
		log.Println("error validasi data", err.Error())
		return todo.ToDo{}, err
	}

	result, err := s.m.AddTask(decodeHP, newTodo)
	if err != nil {
		return todo.ToDo{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

func (s *services) UpdateTask(token *jwt.Token, todoID string, updateTask todo.ToDo) (todo.ToDo, error) {
	decodeHP := middlewares.DecodeToken(token)

	var validateUpdate todo.UpdateTodo
	validateUpdate.TaskName = updateTask.TaskName
	validateUpdate.DueDate = updateTask.DueDate
	validateUpdate.Description = updateTask.Description
	err := s.v.Struct(&validateUpdate)
	if err != nil {
		log.Println("error validasi data", err.Error())
		return todo.ToDo{}, err
	}

	todoid, _ := strconv.Atoi(todoID)

	result, err := s.m.UpdateTask(decodeHP, uint(todoid), updateTask)
	if err != nil {
		return todo.ToDo{}, err
	}
	return result, nil
}

func (s *services) SeeAllMyTask(token *jwt.Token) ([]todo.ToDo, error) {
	decodeHP := middlewares.DecodeToken(token)

	result, err := s.m.SeeAllMyTask(decodeHP)
	if err != nil {
		return nil, err
	}
	return result, nil

}
