package user

import (
	"gorm.io/gorm"
)

type User struct {
	HP       string `json:"hp" form:"hp" validate:"required,min=10,max=13,number" gorm:"primaryKey"`
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Password string `json:"-" validate:"required,alphanum"`
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

func (um *UserModel) UpdateUser(hp string, data User) error {
	if err := um.Connection.Model(&data).Where("hp = ?", hp).Update("name", data.Name).Update("password", data.Password).Error; err != nil {
		return err
	}
	return nil
}

func (um *UserModel) GetAllUsers() ([]User, error) {
	var result []User
	if err := um.Connection.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (um *UserModel) MyProfile(hp string) (User, error) {
	var result User
	if err := um.Connection.Where("hp = ?", hp).First(&result).Error; err != nil {
		return User{}, err
	}
	return result, nil
}

func (um *UserModel) CheckUser(hp string) bool {
	var data User
	if err := um.Connection.Where("hp = ?", hp).First(&data).Error; err != nil {
		return false
	}
	return true
}
