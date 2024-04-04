package main

import (
	"fmt"
	"gin-todoapp/src/config"
	"gin-todoapp/src/models"
	"gin-todoapp/src/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	// Creating a connection to the database
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run()
}