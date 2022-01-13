package usecase

import (
	"ads/domain"
	"ads/internal/infrastructure/database/collection"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateAdCmd func(ctx context.Context, token string, adID string, ad domain.UpdateAd) (domain.Ad, error)

func UpdateAd(client *mongo.Client) UpdateAdCmd {
	return func(ctx context.Context, token string, adID string, ad domain.UpdateAd) (domain.Ad, error) {
		return collection.Update(ctx, token, adID, ad, client)
	}
}
