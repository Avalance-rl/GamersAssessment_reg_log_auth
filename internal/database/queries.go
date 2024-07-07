package database

import (
	"context"
	"dev/reglogauth/internal/models"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(user models.User) (*mongo.InsertOneResult, error) {
	coll := DB.Collection("users")
	doc := models.User{
		Email:            user.Email,
		UserName:         user.UserName,
		Password:         user.Password,
		RegistrationTime: user.RegistrationTime,
	}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		slog.Error("%d", err)
		return nil, err
	}
	return result, err
}
