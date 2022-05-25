package models

import "fmt"

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Grade uint8  `json:"grade"`
}

var Students []Student

func GetStudents() *[]Student {
	return &Students
}

func GetStudentByID(id string) (*Student, error) {
	var s Student
	isExist := false
	for _, student := range Students {
		if student.Id == id {
			s = student
			isExist = true
			break
		}
	}

	if !isExist {
		return nil, fmt.Errorf("no student with id %s", id)
	}

	return &s, nil

}

func AddNewStudent(student *Student) {
	Students = append(Students, *student)
}
