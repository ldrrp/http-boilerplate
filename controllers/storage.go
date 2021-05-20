package controllers

import (
	"encoding/json"
	"http-boilerplate/models"
	"io/ioutil"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

//ArticleController ...
type StorageController struct{}

//Create ...
func (ctrl StorageController) Create(c *gin.Context) {
	jsonBody, errReading := ioutil.ReadAll(c.Request.Body)
	if errReading != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "invalid parameter"})
		return
	}

	var pair Request
	if err := json.Unmarshal(jsonBody, &pair); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "invalid json body"})
		return
	}

	errCreating := models.Storage.Create(pair.Key, pair.Value)

	if errCreating != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "key could not be created"})
		return
	}

	c.JSON(http.StatusOK, pair)
}

//Get ...
func (ctrl StorageController) Get(c *gin.Context) {
	key, ok := c.GetQuery("key")
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "invalid parameters"})
		return
	}

	value, err := models.Storage.Get(key)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "key not found"})
		return
	}

	c.JSON(http.StatusOK, Request{key, value})
}

//Delete ...
func (ctrl StorageController) Delete(c *gin.Context) {
	jsonBody, errReading := ioutil.ReadAll(c.Request.Body)
	if errReading != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "invalid parameter"})
		return
	}

	var pair Request
	if err := json.Unmarshal(jsonBody, &pair); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "invalid json body"})
		return
	}

	errDeleting := models.Storage.Delete(pair.Key)

	if errDeleting != nil {
		c.JSON(http.StatusOK, gin.H{"message": "key not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "key deleted", "data": pair})
}
