package router

import (
	usersController "hexagonal-architecture/app/users/aplication"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Users
	router.POST("/users", usersController.Create)
	router.GET("/users", usersController.Get)

	return router
}
