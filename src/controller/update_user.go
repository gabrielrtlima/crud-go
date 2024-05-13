package controller

import (
	"net/http"
	"strings"

	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	"github.com/gabrielrtlima/crud-go/src/configuration/validation"
	"github.com/gabrielrtlima/crud-go/src/controller/model/request"
	"github.com/gabrielrtlima/crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser Controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest

	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUserServices(userId, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"updateUser controller executed successfully",
		zap.String("journey", "updateUser"))

	c.Status(http.StatusNoContent)
}
