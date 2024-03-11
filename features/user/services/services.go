package services

import (
	"BelajarAPI/features/user"
	"BelajarAPI/helper"
	"BelajarAPI/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	model user.UserModel
	pm    helper.PasswordManager
	v     *validator.Validate
}

func NewService(m user.UserModel) user.UserService {
	return &service{
		model: m,
		pm:    helper.NewPasswordManager(),
		v:     validator.New(),
	}
}

func (s *service) Register(newData user.User) error {
	var registerValidate user.Register
	registerValidate.HP = newData.HP
	registerValidate.Name = newData.Name
	registerValidate.Password = newData.Password
	err := s.v.Struct(&registerValidate)
	if err != nil {
		log.Println("error validasi data", err.Error())
		return err
	}

	newPassword, err := s.pm.HashPassword(newData.Password)
	if err != nil {
		return errors.New(helper.ServiceGeneralError)
	}
	newData.Password = newPassword

	err = s.model.Register(newData)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}
	return nil
}

func (s *service) Login(loginData user.User) (user.User, string, error) {
	var loginValidate user.Login
	loginValidate.HP = loginData.HP
	loginValidate.Password = loginData.Password
	err := s.v.Struct(&loginValidate)
	if err != nil {
		log.Println("error validasi data", err.Error())
		return user.User{}, "", err
	}

	dbData, err := s.model.Login(loginValidate.HP)
	if err != nil {
		return user.User{}, "", err
	}

	err = s.pm.ComparePassword(loginValidate.Password, dbData.Password)
	if err != nil {
		return user.User{}, "", errors.New(helper.UserCredentialError)
	}

	token, err := middlewares.GenerateJWT(dbData.HP)
	if err != nil {
		return user.User{}, "", errors.New(helper.ServiceGeneralError)
	}

	return dbData, token, nil
}

func (s *service) MyProfile(token *jwt.Token) (user.User, error) {
	decodeHP := middlewares.DecodeToken(token)
	result, err := s.model.MyProfile(decodeHP)
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (s *service) UpdateUser(token *jwt.Token, updateData user.User) error {
	decodeHP := middlewares.DecodeToken(token)

	var updateValidate user.UpdateUser
	updateValidate.Name = updateData.Name
	updateValidate.Password = updateData.Password
	err := s.v.Struct(&updateValidate)
	if err != nil {
		log.Println("error validasi data", err.Error())
		return err
	}

	err = s.model.UpdateUser(decodeHP, updateData)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}

	return nil
}

func (s *service) GetAllUsers(token *jwt.Token) ([]user.User, error) {
	decodeHP := middlewares.DecodeToken(token)
	result, err := s.model.GetAllUsers(decodeHP)
	if err != nil {
		return nil, err
	}
	return result, nil
}
