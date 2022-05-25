package models

type Employee struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Departement string `json:"department,omitempty"`
}
