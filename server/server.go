package main

import (
	"log"
	"net/http"
	"server/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost port=5432 user=postgres password=admin dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func main() {
	app := gin.Default()
	db := Init()
	userRepository := &user.UserRepository{
		Db: *db,
	}
	userService := user.UserService{
		UserRepository: userRepository,
	}
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
