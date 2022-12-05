package services

import (
	"context"
	"sesi5/server/users/models"
	"sesi5/server/users/params"
	"sesi5/server/users/repositories"
)

type UserServices struct {
	repo *repositories.UserRepo
}

func NewUserServices(repo *repositories.UserRepo) *UserServices {
	return &UserServices{
		repo: repo,
	}
}

func (u *UserServices) CreateUser(ctx context.Context, req *params.UserCreate) error {
	return u.repo.AddUser(*req.ParseToModel())
}

func (u *UserServices) GetUsers(ctx context.Context) *[]models.User {
	return u.repo.GetUsers()
}
