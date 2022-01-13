package usecase

import (
	"context"
	"epitech/deliveats/user/domain"
	"epitech/deliveats/user/internal/infrastructure/user"
)

type GetUserAddressCmd func(ctx context.Context, userID string) (domain.Address, error)

func GetUserAddress(store user.Store) GetUserAddressCmd {
	return func(ctx context.Context, userID string) (domain.Address, error) {
		user, err := store.Get(ctx, userID)
		if err != nil {
			return domain.Address{}, err
		}
		return user.Address, nil
	}
}
