package collection

import (
	"ads/domain"
	"ads/internal/utils"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(ctx context.Context, ad domain.Ad, client *mongo.Client) error {
	collection := client.Database("GolangProject").Collection("ads")

	//userID, errParse := utils.ParseToken(ad.UserID)

	//if errParse != nil {
	//	return errParse
	//}

	//ad.UserID = userID

	_, err := collection.InsertOne(context.TODO(), ad)
	if err != nil {
		return domain.ErrCreate
	}

	return nil
}

func Get(ctx context.Context, adID string, client *mongo.Client) (domain.Ad, error) {
	collection := client.Database("GolangProject").Collection("ads")

	var result domain.Ad
	err := collection.FindOne(context.TODO(), bson.M{"adid": adID}).Decode(&result)
	if err != nil {
		return domain.Ad{}, domain.ErrAdNotFound
	}
	return result, nil
}

func Delete(ctx context.Context, adID string, token string, client *mongo.Client) error {
	collection := client.Database("GolangProject").Collection("ads")

	userId, errParse := utils.ParseToken(token)

	if errParse != nil {
		return errParse
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"adid": adID, "userid": userId})
	if err != nil {
		return domain.ErrCantDelete
	}

	return nil
}

func Update(ctx context.Context, token string, adID string, ad domain.UpdateAd, client *mongo.Client) (domain.Ad, error) {
	collection := client.Database("GolangProject").Collection("ads")

	userID, errParse := utils.ParseToken(token)

	if errParse != nil {
		return domain.Ad{}, errParse
	}

	var result domain.Ad
	errors := collection.FindOne(context.TODO(), bson.M{"adid": adID}).Decode(&result)
	if errors != nil {
		return domain.Ad{}, domain.ErrAdNotFound
	}

	if result.UserID != userID {
		return domain.Ad{}, domain.ErrUnauthorized
	}

	finalBody, errJson := json.Marshal(ad)

	if errJson != nil {
		return domain.Ad{}, errJson
	}

	var finalmapbody map[string]interface{}

	if errJson = json.Unmarshal(finalBody, &finalmapbody); errJson != nil {
		return domain.Ad{}, errJson
	}

	filter := bson.M{"adid": adID}

	update := bson.M{
		"$set": finalmapbody,
	}

	var resultUpdate domain.Ad

	collection.FindOneAndUpdate(ctx, filter, update).Decode(&resultUpdate)

	return resultUpdate, nil
}

func GetAll(ctx context.Context, token string, userID string, client *mongo.Client) ([]*domain.Ad, error) {
	collection := client.Database("GolangProject").Collection("ads")

	var results []*domain.Ad

	UserID, errParse := utils.ParseToken(token)

	if errParse != nil {
		return nil, errParse
	}

	if userID != UserID {
		return nil, domain.ErrUnauthorized
	}

	cur, err := collection.Find(context.TODO(), bson.D{{"userid", userID}})
	if err != nil {
		return nil, domain.ErrAdNotFound
	}

	for cur.Next(context.TODO()) {
		var elem domain.Ad
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

func GetByKeys(ctx context.Context, keyword string, client *mongo.Client) ([]*domain.Ad, error) {
	collection := client.Database("GolangProject").Collection("ads")

	var results []*domain.Ad

	filter := bson.M{"title": bson.M{"$regex": keyword}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, domain.ErrAdNotFound
	}

	for cur.Next(context.TODO()) {
		var elem domain.Ad
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
