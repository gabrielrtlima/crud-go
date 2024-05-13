package controller

import (
	"net/http"

	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	"github.com/gabrielrtlima/crud-go/src/configuration/validation"
	"github.com/gabrielrtlima/crud-go/src/controller/model/request"
	"github.com/gabrielrtlima/crud-go/src/model"
	"github.com/gabrielrtlima/crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var UserDomainInterface model.UserDomainInterface

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Name,
		userRequest.Password,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
