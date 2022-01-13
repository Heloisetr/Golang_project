package collections

import (
	"account/domain"
	"account/external"
	"account/internal/utils"
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAccount(ctx context.Context, client *mongo.Client, account domain.Account) error {
	collection := client.Database("GolangProject").Collection("account")

	resultCursor, _ := collection.Find(ctx, bson.M{"$or": []bson.M{
		{"email": account.Email},
		{"login": account.Login},
	}})

	var result []bson.M
	resultCursor.All(ctx, &result)

	if len(result) != 0 {
		return domain.ErrEmailAlreadyUsed
	}

	_, err := collection.InsertOne(ctx, account)

	if err != nil {
		return err
	}

	return nil
}

func Login(ctx context.Context, client *mongo.Client, credentials domain.Login) (string, error) {
	collection := client.Database("GolangProject").Collection("account")

	var result domain.Login

	err := collection.FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&result)

	if err != nil {
		return "", domain.ErrEmailNotFound
	}

	if utils.ComparePassword(result.Password, credentials.Password) == false {
		return "", domain.ErrWrongPassword
	}

	token, err := utils.GenerateToken(result.AccountID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetAccount(ctx context.Context, client *mongo.Client, token string, rqAccountID string) (domain.Account, error) {
	collection := client.Database("GolangProject").Collection("account")

	var result domain.Account

	id, errParse := utils.ParseToken(token)

	if errParse != nil {
		return domain.Account{}, errParse
	}

	err := collection.FindOne(ctx, bson.M{"accountid": rqAccountID}).Decode(&result)

	if err != nil {
		return domain.Account{}, err
	}

	if id != rqAccountID {
		return domain.Account{
			AccountID: result.AccountID,
			Email:     result.Email,
			Login:     result.Login,
			Balance:   -1,
		}, nil
	}

	return result, nil
}

func DeleteAccount(ctx context.Context, client *mongo.Client, token string, rqAccountID string) (string, error) {
	collection := client.Database("GolangProject").Collection("account")

	id, errParse := utils.ParseToken(token)

	if errParse != nil {
		return "", errParse
	}

	if id != rqAccountID {
		return "", domain.ErrAccessUnauthorized
	}

	ads, errExt := external.GetAllUserAds(token, id)

	if errExt != nil {
		return "", errExt
	}

	for _, value := range ads {
		errExt = external.DeleteUserAd(token, value.AdID)

		if errExt != nil {
			return "", errExt
		}
	}

	bids, errExt := external.GetAllUserBids(token)

	if errExt != nil {
		return "", errExt
	}

	for _, value := range bids {
		errExt = external.DeleteUserBid(token, value.BidID)

		if errExt != nil {
			return "", errExt
		}
	}

	_, err := collection.DeleteOne(ctx, bson.M{"accountid": id})

	if err != nil {
		return "", domain.ErrAccountNotFound
	}

	return "Account deleted", nil
}

func AddBalance(ctx context.Context, client *mongo.Client, token string, balance float32) (string, error) {
	collection := client.Database("GolangProject").Collection("account")

	id, errParse := utils.ParseToken(token)

	if errParse != nil {
		return "", errParse
	}

	var result domain.Account

	errFind := collection.FindOne(ctx, bson.M{"accountid": id}).Decode(&result)

	if errFind != nil {
		return "", domain.ErrAccountNotFound
	}

	var newBalance = result.Balance + balance

	_, err := collection.UpdateOne(ctx, bson.M{"accountid": id}, bson.D{{Key: "$set", Value: bson.D{{Key: "balance", Value: newBalance}}}})

	if err != nil {
		return "", domain.ErrAccessUnauthorized
	}

	return fmt.Sprintf("New balance of %f", newBalance), nil
}

func UpdateAccount(ctx context.Context, client *mongo.Client, token string, account domain.UpdateAccount) (domain.Account, error) {
	collection := client.Database("GolangProject").Collection("account")

	id, errParse := utils.ParseToken(token)

	if errParse != nil {
		return domain.Account{}, errParse
	}

	finalbody, errJson := json.Marshal(account)

	if errJson != nil {
		return domain.Account{}, errJson
	}

	var finalmapbody map[string]interface{}

	if errJson = json.Unmarshal(finalbody, &finalmapbody); errJson != nil {
		return domain.Account{}, errJson
	}

	filter := bson.M{"accountid": id}

	update := bson.M{
		"$set": finalmapbody,
	}

	var result domain.Account

	collection.FindOneAndUpdate(ctx, filter, update).Decode(&result)

	return result, nil
}
