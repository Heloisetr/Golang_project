package usecase

import (
	"context"
	"transaction/domain"
	"transaction/internal/infrastructure/database/collection"

	"go.mongodb.org/mongo-driver/mongo"
)

type GetAllBidsCmd func(ctx context.Context, token string) ([]*domain.Bid, error)

func GetAllBids(client *mongo.Client) GetAllBidsCmd {
	return func(ctx context.Context, token string) ([]*domain.Bid, error) {
		return collection.GetAll(ctx, token, client)
	}
}
