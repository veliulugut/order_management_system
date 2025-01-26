package database

import (
	"auth-service/pkg/database"
	"context"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (database.User, error)
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
	UpdateUser(ctx context.Context, params database.UpdateUserParams) (database.User, error)
	ListUser(ctx context.Context) ([]database.User, error)
}
