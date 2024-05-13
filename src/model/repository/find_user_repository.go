package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/model"
	"github.com/gabrielrtlima/crud-go/src/model/repository/entity"
	"github.com/gabrielrtlima/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *resterr.RestErr) {
	logger.Info("Init FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
			return nil, resterr.NewNotFoundError(errorMessage)
		}
		errorMessage := fmt.Sprintf("Error trying to find user by email")
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
		return nil, resterr.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed succesfully",
		zap.String("journey", "FindUserByEmail"),
		zap.String("email", email),
		zap.String("user_id", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *resterr.RestErr) {
	logger.Info("Init FindUserByID repository", zap.String("journey", "FindUserByID"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with id: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByID"))
			return nil, resterr.NewNotFoundError(errorMessage)
		}
		errorMessage := fmt.Sprintf("Error trying to find user by id")
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByID"))
		return nil, resterr.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed succesfully",
		zap.String("journey", "FindUserByEmail"),
		zap.String("user_id", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
