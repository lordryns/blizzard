package main

import (
	"blizzard/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello from Blizzard!",
		})
	})

	router.POST("/create_store", controller.CreateRoute)
	router.POST("/add", controller.AddRoute)
	router.GET("/get", controller.GetRoute)
	router.Run()
}
