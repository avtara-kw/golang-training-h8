package main

import "fmt"

type Employee struct {
	Name       string
	Age        int
	Department Department
}

type Department struct {
	Name string
}

func (e *Employee) Print() {
	fmt.Printf("Name :%s \t Age :%d \t Department :%s \n", e.Name, e.Age, e.Department.Name)
}

func main() {
	var dept1 = Department{
		Name: "IT",
	}
	var emp1 Employee = Employee{
		Name:       "Hacktiv8",
		Age:        20,
		Department: dept1,
	}

	var emps = []Employee{
		emp1,
		{
			Name: "Koinworks",
			Age:  15,
			Department: Department{
				Name: "Administration",
			},
		},
	}

	emp1.Name = "Avtara"
	emp1.Department.Name = "Finance"
	emp1.Print()
	for _, emp := range emps {
		emp.Print()
	}

}
