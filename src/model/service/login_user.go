package service

import (
	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *resterr.RestErr) {
	logger.Info("Init loginUser Domain", zap.String("journey", "loginUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if user == nil {
		return nil, "", resterr.NewUnauthorizedError("Email or password is incorrect")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(userDomain.GetPassword()))
	if err != nil {
		return nil, "", resterr.NewUnauthorizedError("Email or password is incorrect")
	}

	token, tokenErr := user.GenerateToken()
	if err != nil {
		return nil, "", tokenErr
	}

	logger.Info("loginUser service executed successfully", zap.String("journey", "loginUser"), zap.String("email", user.GetEmail()))

	return user, token, nil
}
