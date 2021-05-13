package usecase

import (
	"ads/domain"
	"ads/internal/infrastructure/database/collection"
	"context"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateAdCmd func(ctx context.Context, ad domain.Ad) (string, error)

func CreateAd(client *mongo.Client) CreateAdCmd {
	return func(ctx context.Context, ad domain.Ad) (string, error) {
		ad.AdID = uuid.NewV4().String()
		err := collection.Create(ctx, ad, client)
		if err != nil {
			return "", err
		}
		return ad.AdID, nil
	}
}
