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

func AddRoute(ctx *gin.Context) {
	type Payload struct {
		Id    string `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	var payload Payload
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Malformed json, use id, key, value instead.",
			"status":  http.StatusBadRequest,
		})

		return
	}

	var obj = storage.StoreObject{Key: payload.Key, Value: payload.Value}
	var res = store.Add(payload.Id, obj)

	if !res {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to add to store, check id or json content.",
			"status":  http.StatusBadRequest,
		})

		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Object added successfully!",
		"status":  http.StatusOK,
	})

}

func GetRoute(ctx *gin.Context) {
	var id = ctx.Query("id")
	var key = ctx.Query("key")

	if len(id) < 1 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Please specify a store id! hint: /get?id=",
			"status":  http.StatusBadRequest,
		})

		return
	}

	var resp, err = store.Get(id, key)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Store does not exist!",
			"status":  http.StatusBadRequest,
		})

		return
	}

	print(resp)
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Objects pulled successfully!",
		"status":  http.StatusOK,
		"objects": resp,
	})

}

func DeleteRoute(ctx *gin.Context) {
	type Payload struct {
		Id  string `json:"id"`
		Key string `json:"key"`
	}

	var payload Payload
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Malformed json, use id, key instead.",
			"status":  http.StatusBadRequest,
		})

		return
	}

	if !store.Delete(payload.Id, payload.Key) {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "No store with that id!",
			"status":  http.StatusBadRequest,
		})

		return

	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Object deleted successfully!",
		"status":  http.StatusOK,
	})

}
