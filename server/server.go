package main

import (
	"net/http"
	"server/database"
	dtos "server/dtos/user"
	"server/repository"
	service "server/service/user"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	dbFacotry := database.DatabaseFacotory{}
	db := dbFacotry.GetDataBase()
	userRepositoryFactory := &repository.UserRepositoryFactory{}
	userRepository := userRepositoryFactory.CreateUserRespotirory(db)
	userService := service.UserService{
		UserRepository: userRepository,
	}

	app.GET("/isLive", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello From Go",
		})
	})

	app.POST("/user", func(ctx *gin.Context) {
		var body dtos.CreateUser

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		returnValue := userService.CreateUser(&body)

		ctx.JSON(http.StatusOK, returnValue)

	})

	app.Run()
}
