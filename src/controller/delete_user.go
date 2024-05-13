package controller

import (
	"net/http"

	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	"github.com/gabrielrtlima/crud-go/src/configuration/validation"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser Controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "deleteUser"),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"deleteUser controller executed successfully",
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusNoContent)
}
