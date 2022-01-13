package usecase

import (
	"context"
	"transaction/internal/infrastructure/database/collection"

	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteBidCmd func(ctx context.Context, userID string, token string) error

func DeleteBid(client *mongo.Client) DeleteBidCmd {
	return func(ctx context.Context, bidID string, token string) error {
		return collection.Delete(ctx, bidID, token, client)
	}
}
