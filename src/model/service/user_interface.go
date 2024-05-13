package service

import (
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/model"
	"github.com/gabrielrtlima/crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *resterr.RestErr)
	UpdateUserServices(string, model.UserDomainInterface) *resterr.RestErr
	FindUserByIDServices(id string) (model.UserDomainInterface, *resterr.RestErr)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *resterr.RestErr)
	DeleteUserServices(string) *resterr.RestErr
	LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *resterr.RestErr)
}
