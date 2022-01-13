package collection

import (
	"context"
	"transaction/domain"
	"transaction/external"
	"transaction/internal/infrastructure/types"
	"transaction/internal/infrastructure/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(ctx context.Context, bid domain.Bid, token string, client *mongo.Client) error {
	collection := client.Database("GolangProject").Collection("transaction")

	errExt := external.GetAccountCompareBalance(token, bid.UserID, bid.BidPrice)

	bid.Owner = <-external.GetAdInfos(bid.AdId)

	if errExt != nil {
		return errExt
	}

	status := types.Status(bid.Status)
	if status.IsValid() != nil {
		return types.ErrTypeStatusInvalid
	}

	_, err := collection.InsertOne(context.TODO(), bid)
	if err != nil {
		return err
	}

	return nil
}

func Get(ctx context.Context, bidID string, token string, client *mongo.Client) (domain.Bid, error) {
	collection := client.Database("GolangProject").Collection("transaction")

	accountId, errParse := utils.ParseToken(token)

	if errParse != nil {
		return domain.Bid{}, errParse
	}

	var result domain.Bid
	err := collection.FindOne(context.TODO(), bson.M{"bidid": bidID}).Decode(&result)
	if err != nil {
		return domain.Bid{}, domain.ErrBidNotFound
	}

	if result.UserID != accountId && result.Owner != accountId {
		return domain.Bid{}, domain.ErrUnauthorized
	}
	return result, nil
}

func Delete(ctx context.Context, bidID string, token string, client *mongo.Client) error {
	collection := client.Database("GolangProject").Collection("transaction")

	accountId, errParse := utils.ParseToken(token)

	if errParse != nil {
		return errParse
	}

	var result domain.Bid
	err := collection.FindOne(context.TODO(), bson.M{"bidid": bidID}).Decode(&result)
	if err != nil {
		return domain.ErrBidNotFound
	}

	if result.UserID != accountId {
		return domain.ErrUnauthorized
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"bidid": bidID})
	if err != nil {
		return domain.ErrCantDelete
	}

	return nil
}

func Update(ctx context.Context, bidID string, bid domain.Bid, token string, client *mongo.Client) (domain.Bid, error) {
	collection := client.Database("GolangProject").Collection("transaction")

	accountId, errParse := utils.ParseToken(token)

	if errParse != nil {
		return domain.Bid{}, errParse
	}

	var result domain.Bid
	err := collection.FindOne(context.TODO(), bson.M{"bidid": bidID}).Decode(&result)
	if err != nil {
		return domain.Bid{}, domain.ErrBidNotFound
	}

	if result.Owner != accountId {
		return domain.Bid{}, domain.ErrUnauthorized
	}

	status := types.Status(bid.Status)
	if status.IsValid() != nil {
		return domain.Bid{}, types.ErrTypeStatusInvalid
	}

	filter := bson.M{"bidid": bidID}
	update := bson.D{
		{"$set", bson.D{
			{"status", bid.Status},
		},
		},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.Bid{}, domain.ErrCantUpdate
	}

	return result, nil
}

func GetAll(ctx context.Context, token string, client *mongo.Client) ([]*domain.Bid, error) {
	collection := client.Database("GolangProject").Collection("transaction")

	accountId, errParse := utils.ParseToken(token)

	if errParse != nil {
		return nil, errParse
	}

	var results []*domain.Bid

	cur, err := collection.Find(context.TODO(), bson.D{{"userid", accountId}})
	if err != nil {
		return nil, domain.ErrBidNotFound
	}

	for cur.Next(context.TODO()) {
		var elem domain.Bid
		err := cur.Decode((&elem))
		if err != nil {
			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	return results, nil
}
