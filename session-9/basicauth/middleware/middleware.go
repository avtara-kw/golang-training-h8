package middleware

import (
	"encoding/json"
	"net/http"
)

const (
	USERNAME = "batman"
	PASSWORD = "superman"
)

func Auth(rw http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		outputJson(rw, "something went wrong")
		return false
	}

	isValid := (USERNAME == username) && (password == PASSWORD)
	if !isValid {
		outputJson(rw, "username / password is wrong")
		return false
	}
	return true
}

func outputJson(rw http.ResponseWriter, payload interface{}) {
	response := map[string]interface{}{
		"error": payload,
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(401)
	json.NewEncoder(rw).Encode(response)
}
