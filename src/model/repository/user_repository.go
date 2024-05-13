package repository

import (
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *resterr.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *resterr.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *resterr.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *resterr.RestErr
	DeleteUser(userId string) *resterr.RestErr
}
