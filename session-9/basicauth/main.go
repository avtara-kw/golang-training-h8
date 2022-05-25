package main

import (
	"basicauth/middleware"
	"basicauth/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/students", GetStudent)

	models.AddNewStudent(&models.Student{Id: "1", Name: "Hacktiv8", Grade: 1})
	models.AddNewStudent(&models.Student{Id: "2", Name: "Koinworks", Grade: 2})
	models.AddNewStudent(&models.Student{Id: "3", Name: "NooBee", Grade: 1})

	server := new(http.Server)
	port := ":4444"

	fmt.Println("Server running at port", port)
	server.Addr = port
	server.ListenAndServe()
}

func GetStudent(rw http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(rw, r) {
		return
	}

	query := r.URL.Query()
	id := query.Get("id")

	if id == "" {
		outputJson(rw, models.GetStudents())
	} else {
		res, err := models.GetStudentByID(id)
		if err != nil {
			outputJson(rw, err.Error())
			return
		}
		outputJson(rw, res)
	}
}

func outputJson(rw http.ResponseWriter, payload interface{}) {
	response := map[string]interface{}{
		"payload": payload,
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}
