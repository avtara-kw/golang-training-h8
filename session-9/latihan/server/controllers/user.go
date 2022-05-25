package controllers

import (
	"encoding/json"
	"latihan/params"
	"latihan/server/middleware"
	"latihan/services"
	"net/http"
)

type UserController struct {
	userService *services.UserServices
	middleware  *middleware.Middleware
}

func NewUserController(userService *services.UserServices, middleware *middleware.Middleware) *UserController {
	return &UserController{
		userService: userService,
		middleware:  middleware,
	}
}

func (u *UserController) Register(rw http.ResponseWriter, r *http.Request) {
	var req params.UserCreate

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(rw, &response)
		return
	}

	response := u.userService.CreateUser(&req)

	params.WriteJsonResponse(rw, response)
}

func (u *UserController) GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	if !u.middleware.CheckAuth(rw, r) {
		return
	}
	response := u.userService.GetAllUsers()
	params.WriteJsonResponse(rw, response)
}
