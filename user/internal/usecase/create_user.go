package usecase

import (
	"context"
	"epitech/deliveats/user/domain"
	"epitech/deliveats/user/internal/infrastructure/user"

	uuid "github.com/satori/go.uuid"
)

type CreateUserCmd func(ctx context.Context, user domain.User) (string, error)

func CreateUser(store user.Store) CreateUserCmd {
	return func(ctx context.Context, user domain.User) (string, error) {
		user.UserID = uuid.NewV4().String()

		err := store.Create(ctx, user)
		if err != nil {
			return "", err
		}
		return user.UserID, nil
	}
}
