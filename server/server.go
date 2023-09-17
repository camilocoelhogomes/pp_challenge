package main

import (
	"net/http"
	"server/user"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	userService := user.UserService{}
	app.GET("/isLive", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello From Go",
		})
	})

	app.POST("/user", func(ctx *gin.Context) {
		var body user.CreateUser

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
