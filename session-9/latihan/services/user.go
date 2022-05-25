package services

import (
	"latihan/models"
	"latihan/params"
	"latihan/repositories"
	"net/http"
)

type UserServices struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *UserServices {
	return &UserServices{
		userRepo: userRepo,
	}
}

func (u *UserServices) CreateUser(req *params.UserCreate) *params.Response {
	user := makeModelFromRequest(req)
	err := u.userRepo.CreateUser(user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "CREATED SUCCESS",
	}
}

func (u *UserServices) GetAllUsers() *params.Response {
	users, _ := u.userRepo.GetAllUsers()
	return &params.Response{
		Status:  http.StatusOK,
		Payload: users,
	}

}

func makeModelFromRequest(req *params.UserCreate) *models.User {
	return &models.User{
		Username: req.Username,
		Password: req.Password,
	}
}
