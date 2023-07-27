package aplication

import (
	userServices "hexagonal-architecture/app/users/domain/services"
	userDatabase "hexagonal-architecture/app/users/infraestructure"

	// Replace for dto
	userModel "hexagonal-architecture/app/users/domain/models"

	"github.com/gin-gonic/gin"
)

// Create a UserDb instance
var userDb = userDatabase.NewUserDb()

// Create a Service instance using the UserDb
var service = userServices.NewService(userDb)

func Create(c *gin.Context) {
	var user *userModel.User
	c.ShouldBindJSON(&user)

	c.JSON(200, gin.H{
		"data": service.Create(user),
	})
}

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": service.Get(),
	})
}
