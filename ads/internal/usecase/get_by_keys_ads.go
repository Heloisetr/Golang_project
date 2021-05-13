package usecase

import (
	"ads/domain"
	"ads/internal/infrastructure/database/collection"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type GetByKeysAdCmd func(ctx context.Context, keyword string) ([]*domain.Ad, error)

func GetByKeysAd(client *mongo.Client) GetByKeysAdCmd {
	return func(ctx context.Context, keyword string) ([]*domain.Ad, error) {
		return collection.GetByKeys(ctx, keyword, client)
	}
}
