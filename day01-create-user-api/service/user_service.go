package service

import (
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/models"
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/repository"
	"context"
	"errors"
)

type UserService struct {
	Repo *repository.UserReppository
}

func NewUserService(repo *repository.UserReppository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	return s.Repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserById(ctx context.Context, id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}
	return s.Repo.GetUserById(ctx, id)
}
func (s *UserService) DeleteUserBYId(ctx context.Context, id int) error {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}
	return s.Repo.DeleteUserBYId(ctx, id)
}
