package main

import (
	"assignment-3/controllers"
	"assignment-3/models"
	"assignment-3/repositories"
	"fmt"
	"net/http"
)

func main() {
	data := models.Weather{}
	go repositories.GetData(&data)
	http.HandleFunc("/", controllers.AutoReloadWeb)
	fmt.Println("listening on PORT:", ":8080")
	http.ListenAndServe(":8080", nil)
}
