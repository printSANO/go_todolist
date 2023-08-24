package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/printSANO/go_todolist/my-api/Config"
	"github.com/printSANO/go_todolist/my-api/Models"
	"github.com/printSANO/go_todolist/my-api/Routes"
)

var err error

func main() {

	Config.DB, err = gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
