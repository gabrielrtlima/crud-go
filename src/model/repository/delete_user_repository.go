package repository

import (
	"context"
	"os"

	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *resterr.RestErr {
	logger.Info("Init deleteUser repository", zap.String("journey", "deleteUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		errorMessage := "Error trying to convert userId to ObjectID"
		logger.Error(errorMessage, err, zap.String("journey", "deleteUser"))
		return resterr.NewBadRequestError(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		errorMessage := "Error trying to delete user"
		logger.Error(errorMessage, err, zap.String("journey", "deleteUser"))
		return resterr.NewInternalServerError(err.Error())
	}

	return nil
}
