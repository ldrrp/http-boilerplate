package main

import (
	"net/http"

	"github.com/gin-contrib/gzip"

	"http-boilerplate/controllers"
	"http-boilerplate/models"

	"github.com/gin-gonic/gin"
)

var storageModel = make(models.StorageModel)

func main() {
	r := gin.Default()

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	/*** START Storage ***/
	models.Storage = make(models.StorageModel)
	storage := new(controllers.StorageController)

	r.POST("/set", storage.Create)
	r.GET("/get", storage.Get)
	r.POST("/delete", storage.Delete)

	r.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"}) })

	// Start and run the server
	r.Run(":4000")
}
