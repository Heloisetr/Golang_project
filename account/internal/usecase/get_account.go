package usecase

import (
	"account/domain"
	"account/internal/infrastructure/db/collections"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type GetAccountCmd func(ctx context.Context, token string, rqAccountID string) (domain.Account, error)

func GetAccount(client *mongo.Client) GetAccountCmd {
	return func(ctx context.Context, token string, rqAccountID string) (domain.Account, error) {
		account, err := collections.GetAccount(ctx, client, token, rqAccountID)
		if err != nil {
			return domain.Account{}, err
		}
		return account, nil
	}
}
