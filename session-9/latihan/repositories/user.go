package repositories

import (
	"latihan/models"
)

type UserRepo interface {
	CreateUser(model *models.User) error
	GetAllUsers() (*[]models.User, error)
}

type userRepo struct {
	users *[]models.User
}

func NewUserRepo(users *[]models.User) UserRepo {
	return &userRepo{
		users: users,
	}
}

func (u *userRepo) CreateUser(model *models.User) error {
	*u.users = append(*u.users, *model)

	return nil
}

func (u *userRepo) GetAllUsers() (*[]models.User, error) {
	return u.users, nil
}
