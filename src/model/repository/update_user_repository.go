package repository

import (
	"context"
	"os"

	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/model"
	"github.com/gabrielrtlima/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *resterr.RestErr {
	logger.Info("Init updateUser repository", zap.String("journey", "updateUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		errorMessage := "Error trying to convert userId to ObjectID"
		logger.Error(errorMessage, err, zap.String("journey", "updateUser"))
		return resterr.NewBadRequestError(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		errorMessage := "Error trying to update user"
		logger.Error(errorMessage, err, zap.String("journey", "updateUser"))
		return resterr.NewInternalServerError(err.Error())
	}

	return nil
}
