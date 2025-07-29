package main

import (
	"blizzard/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
	}))
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello from Blizzard!",
		})
	})

	router.POST("/create_store", controller.CreateRoute)
	router.POST("/add", controller.AddRoute)
	router.GET("/get", controller.GetRoute)
	router.POST("/delete", controller.DeleteRoute)
	router.Run()
}
