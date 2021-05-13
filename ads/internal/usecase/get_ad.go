package usecase

import (
	"ads/domain"
	"ads/internal/infrastructure/database/collection"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type GetAdCmd func(ctx context.Context, adID string) (domain.Ad, error)

func GetAd(client *mongo.Client) GetAdCmd {
	return func(ctx context.Context, adID string) (domain.Ad, error) {
		return collection.Get(ctx, adID, client)
	}
}
