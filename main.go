package main

import (
	"context"
	"log"

	mongodb "github.com/gabrielrtlima/crud-go/src/configuration/database"
	"github.com/gabrielrtlima/crud-go/src/configuration/logger"
	"github.com/gabrielrtlima/crud-go/src/controller"
	"github.com/gabrielrtlima/crud-go/src/controller/routes"
	"github.com/gabrielrtlima/crud-go/src/model/repository"
	"github.com/gabrielrtlima/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("About to start user application")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
