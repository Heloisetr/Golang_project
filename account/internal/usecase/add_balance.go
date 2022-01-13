package usecase

import (
	"account/internal/infrastructure/db/collections"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AddBalanceCmd func(ctx context.Context, token string, balance float32) (string, error)

func AddBalance(client *mongo.Client) AddBalanceCmd {
	return func(ctx context.Context, token string, balance float32) (string, error) {
		message, err := collections.AddBalance(ctx, client, token, balance)

		if err != nil {
			return "", err
		}

		return message, nil
	}
}
