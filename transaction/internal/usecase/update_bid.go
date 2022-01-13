package usecase

import (
	"context"
	"transaction/domain"
	"transaction/internal/infrastructure/database/collection"

	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateBidCmd func(ctx context.Context, bidID string, bid domain.Bid, token string) (domain.Bid, error)

func UpdateBid(client *mongo.Client) UpdateBidCmd {
	return func(ctx context.Context, bidID string, bid domain.Bid, token string) (domain.Bid, error) {
		return collection.Update(ctx, bidID, bid, token, client)
	}
}
