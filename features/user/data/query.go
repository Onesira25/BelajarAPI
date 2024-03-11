package data

import (
	"BelajarAPI/features/user"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) user.UserModel {
	return &model{
		connection: db,
	}
}

func (m *model) Register(newData user.User) error {
	err := m.connection.Create(&newData).Error

	if err != nil {
		return errors.New("terjadi masalah pada database")
	}
	return nil
}

func (m *model) Login(hp string) (user.User, error) {
	var result user.User
	if err := m.connection.Model(&User{}).Where("hp = ? ", hp).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (m *model) UpdateUser(hp string, data user.User) error {
	if err := m.connection.Model(&data).Where("hp = ?", hp).Update("name", data.Name).Update("password", data.Password).Error; err != nil {
		return err
	}

	return nil
}

func (m *model) GetAllUsers(hp string) ([]user.User, error) {
	var result []user.User

	if err := m.connection.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (m *model) MyProfile(hp string) (user.User, error) {
	var result user.User
	if err := m.connection.Model(&User{}).Where("hp = ?", hp).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (m *model) CheckUser(hp string) bool {
	var data user.User
	if err := m.connection.Where("hp = ?", hp).First(&data).Error; err != nil {
		return false
	}
	return true
}
