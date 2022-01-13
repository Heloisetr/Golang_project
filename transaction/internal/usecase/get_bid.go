package usecase

import (
	"context"
	"transaction/domain"
	"transaction/internal/infrastructure/database/collection"

	"go.mongodb.org/mongo-driver/mongo"
)

type GetBidCmd func(ctx context.Context, bidID string, token string) (domain.Bid, error)

func GetBid(client *mongo.Client) GetBidCmd {
	return func(ctx context.Context, bidID string, token string) (domain.Bid, error) {
		return collection.Get(ctx, bidID, token, client)
	}
}
