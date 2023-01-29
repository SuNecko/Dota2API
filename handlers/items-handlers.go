package handlers

import (
	"ApiCRUD/db"
	"ApiCRUD/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetItemsHandler(c *gin.Context) {
	var items []models.Item
	db.DB.Find(&items)

	encodeErr := json.NewEncoder(c.Writer).Encode(&items)
	if encodeErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}
}

func GetItemHandler(c *gin.Context) {
	var item models.Item
	name := c.Param("name")

	db.DB.Where("name=?", name).Find(&item)

	if item.Id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Error": "Item not Found",
		})
	}

}

func PostItemHandler(c *gin.Context) {
	var item models.Item

	decodeErr := json.NewDecoder(c.Request.Body).Decode(&item)
	if decodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": decodeErr})
	}

	createdItem := db.DB.Create(&item)
	err := createdItem.Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "DB Error",
		})
	}

	c.Writer.WriteHeader(http.StatusCreated)
	encodeErr := json.NewEncoder(c.Writer).Encode(&item)
	if encodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}

}

func DeleteItemHandler(c *gin.Context) {
	var item models.Item
	id := c.Param("id")

	db.DB.Find(&item, id).Delete(&item)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successful",
		"Value":   item,
	})
}
