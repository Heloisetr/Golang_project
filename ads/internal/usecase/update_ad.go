package usecase

import (
	"ads/domain"
	"ads/internal/infrastructure/database/collection"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateAdCmd func(ctx context.Context, adID string, ad domain.Ad) (domain.Ad, error)

func UpdateAd(client *mongo.Client) UpdateAdCmd {
	return func(ctx context.Context, adID string, ad domain.Ad) (domain.Ad, error) {
		return collection.Update(ctx, adID, ad, client)
	}
}
