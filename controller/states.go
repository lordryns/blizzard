package controller

import (
	"blizzard/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var store = storage.NewStorage()

func CreateRoute(ctx *gin.Context) {
	type Payload struct {
	}

	var payload Payload

	if err := ctx.Bind(&payload); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong!",
			"status":  http.StatusBadRequest,
		})
		return
	}

	id := uuid.New().String()
	var isCreated = store.Create(id)
	if !isCreated {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Unable to create store, please try again!",
			"status":  http.StatusBadRequest,
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Store created successfully!",
		"status":  http.StatusOK,
		"id":      id,
	})

}
