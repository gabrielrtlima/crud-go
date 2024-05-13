package controller

import (
	"net/http"
	"net/mail"

	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gabrielrtlima/crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID Controller",
		zap.String("journey", "findUserByID"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := resterr.NewBadRequestError("User ID is not valid.")
		logger.Error("Error trying to validate userId", err,
			zap.String("journey", "findUserByID"),
			zap.String("userId", userId),
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID services", err,
			zap.String("journey", "findUserByID"),
			zap.String("userId", userId),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Success to call findUserByID services",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userId))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	// code to find user by email
	logger.Info("Init findUserByEmail Controller",
		zap.String("journey", "findUserByEmail"),
	)

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := resterr.NewBadRequestError("User email is not valid.")
		logger.Error("Error trying to validate userEmail", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail),
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail services", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Success to call findUserByEmail services",
		zap.String("journey", "findUserByEmail"),
		zap.String("userEmail", userEmail))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
