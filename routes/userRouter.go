package routes

import (
	"github.com/Kiteretzu/gin-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.UserController.GetUsers())
	incomingRoutes.GET("/users/:id", controllers.UserController.GetUser())
}