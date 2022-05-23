package main

import (
	"fmt"
	"sesi7-gorm/database"
	"sesi7-gorm/models"
	"sesi7-gorm/repository"
	"strings"
)

func main() {
	db := database.StartDB()
	userRepo := repository.NewUserRepo(db)

	user := models.User{
		Email: "test@gmail.com",
	}

	err := userRepo.CreateUser(&user)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	fmt.Println("Created success")

	employees, err := userRepo.GetAllUsers()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for k, emp := range *employees {
		fmt.Println("User :", k+1)
		emp.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

	// emp, err := userRepo.GetUserByID(3)

	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// emp.Print()
}
