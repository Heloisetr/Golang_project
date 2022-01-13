package user

import (
	"context"
	"epitech/deliveats/user/domain"
)

type Store interface {
	Create(ctx context.Context, user domain.User) error
	Get(ctx context.Context, userID string) (domain.User, error)
	Delete(ctx context.Context, userID string) error
}
