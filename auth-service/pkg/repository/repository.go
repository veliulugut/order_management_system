package database

import (
	"auth-service/pkg/database"
	"context"
	"errors"
)

var _ UserRepositoryInterface = (*Repository)(nil)

func NewRepository(queries *database.Queries) *Repository {
	return &Repository{
		queries: queries,
	}
}

type Repository struct {
	queries *database.Queries
}

// CreateUser implements UserRepositoryInterface.
func (r *Repository) CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error) {
	users, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return database.User{}, errors.New("failed to create user")
	}

	return users, nil
}

// DeleteUser implements UserRepositoryInterface.
func (r *Repository) DeleteUser(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetUser implements UserRepositoryInterface.
func (r *Repository) GetUser(ctx context.Context, id int64) (database.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements UserRepositoryInterface.
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserRepositoryInterface.
func (r *Repository) UpdateUser(ctx context.Context, user database.UpdateUserParams) (database.User, error) {
	panic("unimplemented")
}

func (r *Repository) ListUser(ctx context.Context) ([]database.User, error) {
	panic("unimplemented")
}
