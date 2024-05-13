package service

import (
	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(userId string, userDomain model.UserDomainInterface) *resterr.RestErr {
	logger.Info("Init updateUser Domain", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		return err
	}

	return nil
}
