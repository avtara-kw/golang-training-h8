package main
import (
	"session-3/params"
	"session-3/repository"
)

func main(){
	reqs := []params.User{
		params.User{Name: "Joko",Email: "jokosuracil112@gmail.com"},
		params.User{Name: "Joko",Email: "jokosuracil112@gmail.com"},
	}

	for _, s := range reqs {
		repository.InsertUser(&s)
	}

	repository.ReadUser()

	// repository.InsertUser
}