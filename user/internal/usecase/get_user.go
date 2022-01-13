package usecase

import (
	"context"
	"epitech/deliveats/user/domain"
	"epitech/deliveats/user/internal/infrastructure/user"
)

type GetUserCmd func(ctx context.Context, userID string) (domain.User, error)

func GetUser(store user.Store) GetUserCmd {
	return func(ctx context.Context, userID string) (domain.User, error) {
		return store.Get(ctx, userID)
	}
}
