package usecase

import (
	"account/internal/infrastructure/db/collections"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteAccountCmd func(ctx context.Context, token string, rqAccountID string) (string, error)

func DeleteAccount(client *mongo.Client) DeleteAccountCmd {
	return func(ctx context.Context, token string, rqAccountID string) (string, error) {
		message, err := collections.DeleteAccount(ctx, client, token, rqAccountID)
		if err != nil {
			return "", err
		}
		return message, nil
	}
}
