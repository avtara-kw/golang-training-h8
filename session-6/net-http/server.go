package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sesi6/data"
	"sesi6/models"
	"sesi6/params"
)

func main() {
	port := ":9999"

	http.HandleFunc("/", hello)
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employees/create", createEmployee)

	fmt.Println("Server running at port ", port)
	http.ListenAndServe(port, nil)
}

func greet(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello guys"))
}

func hello(rw http.ResponseWriter, r *http.Request) {
	student := map[string]interface{}{
		"name": "Hacktiv8",
		"age":  20,
	}

	studentByte, _ := json.Marshal(student)
	rw.Write(studentByte)
}

func getEmployees(rw http.ResponseWriter, r *http.Request) {

	clientId := r.Header.Get("X-CLIENT-ID")

	method := r.Method

	if method == http.MethodGet {
		query := r.URL.Query()
		nameSearch := query.Get("name")

		if clientId == "API" {
			rw.Header().Add("Content-Type", "application/json")
			if nameSearch == "" {
				json.NewEncoder(rw).Encode(data.Employee)
			} else {
				emps := data.Employee
				var emp models.Employee

				for _, e := range emps {
					if e.Name == nameSearch {
						emp = e
						break
					}
				}

				if emp.Name == "" {
					rw.WriteHeader(http.StatusNotFound)
					json.NewEncoder(rw).Encode(map[string]interface{}{
						"error": 404,
						"msg":   "DATA NOT FOUND",
					})
					return
				}

				json.NewEncoder(rw).Encode(emp)
			}
		} else {
			tpl, err := template.ParseFiles("./static/index.html")
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			tpl.Execute(rw, data.Employee)
			return
		}

	} else {
		httpMethodNotAllowed(rw)
	}
}

func createEmployee(rw http.ResponseWriter, r *http.Request) {
	method := r.Method
	rw.Header().Add("Content-Type", "application/json")
	if method == http.MethodPost {
		var request params.CreateEmployee

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"error":           http.StatusBadRequest,
				"msg":             "BAD REQUEST",
				"additional_info": err.Error(),
			})
			return
		}

		emp := models.Employee{
			Name:        request.Name,
			Age:         request.Age,
			Departement: request.Departement,
		}

		data.Employee = append(data.Employee, emp)

		json.NewEncoder(rw).Encode(data.Employee)
	} else {
		httpMethodNotAllowed(rw)
	}
}

func httpMethodNotAllowed(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"error": 405,
		"msg":   "METHOD NOT ALLOWED",
	})
}
