package usecase

import (
	"account/domain"
	"account/internal/infrastructure/db/collections"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoginCmd func(ctx context.Context, credentials domain.Login) (string, error)

func Login(client *mongo.Client) LoginCmd {
	return func(ctx context.Context, credentials domain.Login) (string, error) {
		token, err := collections.Login(ctx, client, credentials)

		if err != nil {
			return "", err
		}

		return token, nil
	}
}
