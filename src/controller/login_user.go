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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser Controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"LoginUser controller executed successfully",
		zap.String("journey", "createUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
