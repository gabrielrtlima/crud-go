package service

import (
	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *resterr.RestErr) {
	logger.Info("Init CreateUser Domain", zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		return nil, resterr.NewBadRequestError("User already exists")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
