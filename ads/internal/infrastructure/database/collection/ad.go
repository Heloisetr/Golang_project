package collection

import (
	"ads/domain"
	"ads/internal/utils"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(ctx context.Context, ad domain.Ad, client *mongo.Client) error {
	collection := client.Database("GolangAd").Collection("ads")

	_, err := collection.InsertOne(context.TODO(), ad)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func Get(ctx context.Context, adID string, client *mongo.Client) (domain.Ad, error) {
	collection := client.Database("GolangAd").Collection("ads")

	var result domain.Ad
	err := collection.FindOne(context.TODO(), bson.M{"adid": adID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func Delete(ctx context.Context, adID string, client *mongo.Client) error {
	collection := client.Database("GolangAd").Collection("ads")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"adid": adID})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func Update(ctx context.Context, adID string, ad domain.Ad, client *mongo.Client) (domain.Ad, error) {
	collection := client.Database("GolangAd").Collection("ads")

	var result domain.Ad
	errors := collection.FindOne(context.TODO(), bson.M{"adid": adID}).Decode(&result)
	if errors != nil {
		log.Fatal(errors)
	}

	filter := bson.M{"adid": adID}
	update := bson.D{
		{"$set", bson.D{
			{"title", utils.CheckEmptyString(result.Title, ad.Title)},
			{"description", utils.CheckEmptyString(result.Description, ad.Description)},
			{"price", utils.CheckEmptyNumber(result.Price, ad.Price)},
			{"picture", utils.CheckEmptyPicture(result.Picture, ad.Picture)},
		},
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func GetAll(ctx context.Context, userID string, client *mongo.Client) ([]*domain.Ad, error) {
	collection := client.Database("GolangAd").Collection("ads")

	var results []*domain.Ad

	cur, err := collection.Find(context.TODO(), bson.D{{"userid", userID}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem domain.Ad
		err := cur.Decode((&elem))
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetByKeys(ctx context.Context, keyword string, client *mongo.Client) ([]*domain.Ad, error) {
	collection := client.Database("GolangAd").Collection("ads")

	var results []*domain.Ad

	filter := bson.M{"title": bson.M{"$regex": keyword}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem domain.Ad
		err := cur.Decode((&elem))
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return results, nil
}
