package usecase

import (
	"account/domain"
	"account/internal/infrastructure/db/collections"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateAccountCmd func(ctx context.Context, token string, update domain.UpdateAccount) (domain.Account, error)

func UpdateAccount(client *mongo.Client) UpdateAccountCmd {
	return func(ctx context.Context, token string, update domain.UpdateAccount) (domain.Account, error) {
		account, err := collections.UpdateAccount(ctx, client, token, update)

		if err != nil {
			return domain.Account{}, err
		}

		return account, nil
	}
}
