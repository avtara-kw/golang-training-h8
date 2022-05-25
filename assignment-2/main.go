package main

import (
	"assignment-2/app/config"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	DB_HOST  = "localhost"
	DB_PORT  = "5432"
	DB_USER  = "docker"
	DB_PASS  = "docker"
	DB_NAME  = "docker"
	APP_PORT = ":8888"
)

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	log.Println("server running at port ", APP_PORT)
	router.Run(APP_PORT)
}
