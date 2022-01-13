package usecase

import (
	"account/domain"
	"account/internal/infrastructure/db/collections"
	"context"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateAccountCmd func(ctx context.Context, account domain.Account) (string, error)

func CreateAccount(client *mongo.Client) CreateAccountCmd {
	return func(ctx context.Context, account domain.Account) (string, error) {
		account.AccountID = uuid.NewV4().String()

		err := collections.CreateAccount(ctx, client, account)
		if err != nil {
			return "", err
		}
		return account.AccountID, nil
	}
}
