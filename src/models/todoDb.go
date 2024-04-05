package models

import (
	"gin-todoapp/src/config"
	_ "github.com/go-sql-driver/mysql"
)

// fetch all todos at once
func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err := config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}

	if err = config.DB.Model(&Todo{}).Where("id = ?", id).Updates(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteATodo(todo *Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
