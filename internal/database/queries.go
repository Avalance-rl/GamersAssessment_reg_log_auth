package database

import (
	"context"
	"dev/reglogauth/internal/models"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(user models.User) error {
	coll := DB.Collection("users")
	doc := models.User{
		Email:            user.Email,
		UserName:         user.UserName,
		Password:         user.Password,
		RegistrationTime: user.RegistrationTime,
	}
	_, err := coll.InsertOne(context.TODO(), doc)
	if mongo.IsDuplicateKeyError(err) {
		slog.Error("Duplicate key error collection")
	}

	return err
}

func InsertToken(email string, token string) string {
	coll := DB.Collection("tokens")
	doc := models.Token{
		Token:        token,
		Email:        email,
		CreationTime: time.Now(),
	}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		slog.Error("insert token", err)
	}
	insertedId := fmt.Sprintf("%v", result.InsertedID)

	return insertedId
}

func FindUser(email string) string {
	coll := DB.Collection("users")
	var result bson.M
	err := coll.FindOne(
		context.TODO(), bson.D{{"email", email}},
	).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		slog.Error("No document was found with the email %s", email)
		return ""
	}
	pass := result["password"].(string)

	return pass
}

func FindToken(tokenID string) bson.M {
	coll := DB.Collection("tokens")
	var result bson.M
	err := coll.FindOne(
		context.TODO(), bson.D{{"_id", tokenID}},
	).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		slog.Error(
			"No document was found with the tokenID %s,"+
				" maybe the token is compromised",
			tokenID)
		return nil
	}

	return result

}
