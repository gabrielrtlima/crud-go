package service

import (
	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserServices(userId string) *resterr.RestErr {
	logger.Info("Init deleteUser Domain", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		return err
	}

	return nil
}
