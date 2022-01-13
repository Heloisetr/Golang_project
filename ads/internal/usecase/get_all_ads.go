package usecase

import (
	"ads/domain"
	"ads/internal/infrastructure/database/collection"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type GetAllAdCmd func(ctx context.Context, token string, userID string) ([]*domain.Ad, error)

func GetAllAd(client *mongo.Client) GetAllAdCmd {
	return func(ctx context.Context, token string, userID string) ([]*domain.Ad, error) {
		return collection.GetAll(ctx, token, userID, client)
	}
}
