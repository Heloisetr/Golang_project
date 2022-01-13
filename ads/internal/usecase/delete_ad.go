package usecase

import (
	"ads/internal/infrastructure/database/collection"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteAdCmd func(ctx context.Context, token string, adID string) error

func DeleteAd(client *mongo.Client) DeleteAdCmd {
	return func(ctx context.Context, token string, adID string) error {
		return collection.Delete(ctx, adID, token, client)
	}
}
