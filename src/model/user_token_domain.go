package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

var JWT_SECRET_KEY = "JWT_SECRET_KEY"

func (ud *userDomain) GenerateToken() (string, *resterr.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", resterr.NewInternalServerError(fmt.Sprintf("Error generating token: %s", err.Error()))
	}

	return tokenString, nil
}

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(c.GetHeader("Authorization"))

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, resterr.NewBadRequestError("Invalid token")
	})
	if err != nil {
		errRest := resterr.NewUnauthorizedError("Invalid token")

		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := resterr.NewUnauthorizedError("Invalid token")

		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}

	logger.Info(fmt.Sprintf("User authenticated %v", userDomain), zap.String("journey", "VerifyTokenMiddleware"))
}

func VerifyToken(tokenValue string) (UserDomainInterface, *resterr.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, resterr.NewBadRequestError("Invalid token")
	})
	if err != nil {
		return nil, resterr.NewUnauthorizedError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, resterr.NewUnauthorizedError("Invalid token")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return token[7:]
	}

	return token
}
