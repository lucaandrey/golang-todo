package models

import (
	"gin-todoapp/src/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(u *[]User) error {
	if err := config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(u *User) error {
	if err := config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetAnUser(id string, u *User) error {
	if err := config.DB.Where("id = ?", id).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(id string, u *User) error {
	if err := config.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}

	if err := config.DB.Model(User{}).Where("id = ?", id).Updates(u).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string, u *User) error {
	if err := config.DB.Where("id = ?", id).Delete(u).Error; err != nil {
		return err
	}
	return nil
}
