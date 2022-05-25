package middleware

import (
	"latihan/models"
	"latihan/params"
	"net/http"
)

type Middleware struct {
	users *[]models.User
}

func NewMiddleware(users *[]models.User) *Middleware {
	return &Middleware{
		users: users,
	}
}

func (m *Middleware) CheckAuth(rw http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		params.WriteJsonResponse(rw, &params.Response{
			Status: 403,
			Error:  "FORBIDDEN ACCESS",
		})
		return false
	}

	isValid := false

	for _, user := range *m.users {
		if user.Username == username && password == user.Password {
			isValid = true
			break
		}
	}

	if !isValid {
		params.WriteJsonResponse(rw, &params.Response{
			Status:         401,
			Error:          "UNAUTHORIZE",
			AdditionalInfo: "username / password is wrong",
		})
		return false
	}

	return true
}
