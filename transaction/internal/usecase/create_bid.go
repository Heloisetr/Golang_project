package usecase

import (
	"context"
	"transaction/domain"
	"transaction/internal/infrastructure/database/collection"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateBidCmd func(ctx context.Context, bid domain.Bid, token string) (string, error)

func CreateBid(client *mongo.Client) CreateBidCmd {
	return func(ctx context.Context, bid domain.Bid, token string) (string, error) {
		bid.BidID = uuid.NewV4().String()
		err := collection.Create(ctx, bid, token, client)
		if err != nil {
			return "", err
		}
		return bid.BidID, nil
	}
}
