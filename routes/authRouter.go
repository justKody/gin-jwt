package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.UserController.SignUp())
	incomingRoutes.POST("/users/login", controllers.UserController.Login())
}
