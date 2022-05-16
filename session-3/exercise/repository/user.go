package repository
import (
	"time"
	"session-3/params"
	"fmt"
)

var store []ModelUser

type ModelUser struct{
	Name string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InsertUser(u *params.User){
	store = append(store, ModelUser{
		Name: u.Name,
		Email: u.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}) 
}

func ReadUser(){
	for i, v := range store{
		fmt.Printf("User %d\n", i)
		fmt.Printf("Name :%v\n", v.Name)
		fmt.Printf("Name :%v\n", v.Email)
		fmt.Printf("Name :%v\n", v.CreatedAt)
		fmt.Printf("Name :%v\n", v.UpdatedAt)
	}
}



