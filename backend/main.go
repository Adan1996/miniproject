package main

import (
	"fmt"

	"example.com/m/Config"
	"example.com/m/Models"
	"example.com/m/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Whislist{})
	r := Routes.SetupRouter()
	r.Run(":2000")
}
