package service

import (
	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) findUser(string) (*model.UserDomainInterface, *resterr.RestErr) {
	return nil, nil
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *resterr.RestErr) {
	logger.Info("Init FindUserByEmail services", zap.String("journey", "FindUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *resterr.RestErr) {
	logger.Info("Init FindUserByID services", zap.String("journey", "FindUserByID"))

	return ud.userRepository.FindUserByID(id)
}
