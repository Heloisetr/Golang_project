package account

import (
	"account/domain"
	"context"
)

type Store interface {
	Create(ctx context.Context, account domain.Account) error
	Get(ctx context.Context, accountID string) (domain.Account, error)
	Login(ctx context.Context, credentials domain.Login) (string, error)
}
