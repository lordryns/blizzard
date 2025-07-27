package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
  "blizzard/controller"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.POST("/create", controller.CreateRoute)
	router.Run()
}

