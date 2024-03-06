package todo

import (
	"BelajarAPI/helper"
	"BelajarAPI/middlewares"
	"BelajarAPI/model/todo"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	Model todo.ToDoModel
}

func (tc *TodoController) AddTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hpFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		var input ToDoRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirimkan tidak sesuai", nil))
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "Data Kurang Sesuai", nil))
		}

		var inputProcess todo.ToDoList
		inputProcess.UserHP = hpFromToken
		inputProcess.TaskName = input.TaskName
		inputProcess.DueDate = input.DueDate
		inputProcess.Description = input.Description

		result, err := tc.Model.AddTask(inputProcess)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan saat update data", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "task added successfully", result))
	}
}

func (tc *TodoController) UpdateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("todoID")
		todoID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		var hpFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		if hpFromToken == "" {
			log.Println("error decode token:", "hp tidak ditemukan")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "tidak dapat mengakses fitur ini", nil))
		}

		var input ToDoUpdate
		err = c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirimkan tidak sesuai", nil))
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "Data Kurang Sesuai", nil))
		}

		var inputProcess todo.ToDoList
		inputProcess.TaskName = input.TaskName
		inputProcess.DueDate = input.DueDate
		inputProcess.Description = input.Description

		updatedTask, err := tc.Model.UpdateTask(hpFromToken, uint(todoID), inputProcess)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan saat update data", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "Task Updated Successfully", updatedTask))
	}
}

func (tc *TodoController) SeeAllMyTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hpFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		task, err := tc.Model.SeeAllTask(hpFromToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "successfully get all task", task))
	}
}
