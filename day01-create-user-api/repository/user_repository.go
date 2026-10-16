package repository

import (
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/models"
	"context"
	"database/sql"

	"github.com/aarondl/sqlboiler/v4/boil"
)

type UserReppository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserReppository {
	return &UserReppository{
		DB: db,
	}
}

func (r *UserReppository) CreateUser(ctx context.Context, user *models.User) error {
	err := user.Insert(ctx, r.DB, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
func (r *UserReppository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	user, err := models.FindUser(ctx, r.DB, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserReppository) DeleteUserBYId(ctx context.Context, id int) error {
	user, err := models.FindUser(ctx, r.DB, id)
	if err != nil {
		return err
	}

	_, err = user.Delete(ctx, r.DB)
	if err != nil {
		return err
	}
	return nil
}
