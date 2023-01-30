package handlers

import (
	"ApiCRUD/db"
	"ApiCRUD/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetHeroesHandler(c *gin.Context) {
	var heroes []models.Hero
	db.DB.Find(&heroes)

	encodeErr := json.NewEncoder(c.Writer).Encode(&heroes)
	if encodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}
}

func GetHeroHandler(c *gin.Context) {
	var hero models.Hero
	name := c.Param("name")

	db.DB.Where("name= ?", name).Find(&hero)
	if hero.Id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Error": "Hero not Found",
		})
	}

	findErr := db.DB.Model(&hero).Association("Skills").Find(&hero.Skills)
	if findErr != nil {
		log.Fatal(findErr)
	}

	encodeErr := json.NewEncoder(c.Writer).Encode(&hero)
	if encodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}
}

func GetHeroByAttributeHandler(c *gin.Context) {
	var heroes []models.Hero
	attribute := c.Param("attribute")

	db.DB.Where("attribute= ?", attribute).Find(&heroes)

	encodeErr := json.NewEncoder(c.Writer).Encode(&heroes)
	if encodeErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}

}

func PostHeroHandler(c *gin.Context) {
	var hero models.Hero

	decodeErr := json.NewDecoder(c.Request.Body).Decode(&hero)
	if decodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": decodeErr})
	}

	createdHero := db.DB.Create(&hero)
	err := createdHero.Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "DB Error",
		})
	}

	c.Writer.WriteHeader(http.StatusCreated)
	encodeErr := json.NewEncoder(c.Writer).Encode(&hero)
	if encodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}

}

func DeleteHeroHandler(c *gin.Context) {
	var hero models.Hero
	id := c.Param("id")

	db.DB.Find(&hero, id).Delete(&hero)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successful",
		"Value":   hero,
	})
}
