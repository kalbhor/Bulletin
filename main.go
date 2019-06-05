package main

import (
	"fmt"
	"./Config"
	"./Models"
	"./Routers"
	"github.com/jinzhu/gorm"
)


var err error

func main() {

	Config.DB, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Category{})
	Config.DB.AutoMigrate(&Models.Service{})
	Config.DB.AutoMigrate(&Models.Offer{})

	r := Routers.SetupRouter()

	r.Run(":8080")
}
