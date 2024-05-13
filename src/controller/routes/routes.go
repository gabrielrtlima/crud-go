package routes

import (
	"github.com/gabrielrtlima/crud-go/src/controller"
	"github.com/gabrielrtlima/crud-go/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PATCH("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
}
