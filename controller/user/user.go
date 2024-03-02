package user

import (
	"BelajarAPI/helper"
	"BelajarAPI/middlewares"
	"BelajarAPI/model"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.UserModel
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterRequest
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

		var processInput model.User
		processInput.HP = input.HP
		processInput.Name = input.Name
		processInput.Password = input.Password

		err = uc.Model.Register(processInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "selamat data sudah terdaftar", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
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

		result, err := uc.Model.Login(input.HP, input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}

		token, err := middlewares.GenerateJWT(result.HP)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem, gagal memproses data", nil))
		}

		var responseData LoginResponse
		responseData.HP = result.HP
		responseData.Name = result.Name
		responseData.Token = token

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "selamat anda berhasil login", responseData))

	}
}

func (uc *UserController) AddTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hpFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))
		var input model.ToDoList
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

		err = uc.Model.AddTask(hpFromToken, input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan saat update data", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "task added successfully", nil))
	}
}

func (uc *UserController) UpdateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hpFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))
		var input model.ToDoList
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

		updatedTask, err := uc.Model.UpdateTask(hpFromToken, input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan saat update data", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "Task Updated Successfully", updatedTask))
	}
}

func (uc *UserController) SeeAllTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		task, err := uc.Model.SeeAllTask()
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "successfully get all task", task))
	}
}
