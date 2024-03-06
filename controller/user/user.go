package user

import (
	"BelajarAPI/helper"
	"BelajarAPI/middlewares"
	"BelajarAPI/model/user"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model user.UserModel
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

		var processInput user.User
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

func (uc *UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hpFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		isFound := uc.Model.CheckUser(hpFromToken)

		if !isFound {
			return c.JSON(http.StatusNotFound,
				helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
		}

		var input UpdateRequest
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

		var processInput user.User
		processInput.HP = hpFromToken
		processInput.Name = input.Name
		processInput.Password = input.Password

		err = uc.Model.UpdateUser(hpFromToken, processInput)

		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan saat update data", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "success update user", nil))
	}
}

func (uc *UserController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.Model.GetAllUsers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data semua user", users))
	}
}

func (uc *UserController) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hpFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

		result, err := uc.Model.MyProfile(hpFromToken)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound,
					helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
			}
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", result))
	}
}
