package main

import (
	"fmt"
	// "log"
	"todo_app/app/models"
	// "todo_app/config"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println(("test"))
	// fmt.Println(models.Db)	

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()
	u, _ := models.GetUser(1)
	fmt.Println(u)
}
