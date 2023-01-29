package handlers

import (
	"ApiCRUD/db"
	"ApiCRUD/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSkillsHandler(c *gin.Context) {
	var skills []models.Skill
	db.DB.Find(&skills)

	encodeErr := json.NewEncoder(c.Writer).Encode(skills)
	if encodeErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}
}

func GetSkillHandler(c *gin.Context) {
	var skill models.Skill
	name := c.Param("name")

	db.DB.Where("name=?", name).Find(&skill)
	if skill.Id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Error": "Item not Found",
		})
	}

	encodeErr := json.NewEncoder(c.Writer).Encode(&skill)
	if encodeErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}
}

func PostSkillsHandler(c *gin.Context) {
	var skill models.Skill

	decodeErr := json.NewDecoder(c.Request.Body).Decode(&skill)
	if decodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": decodeErr})
	}

	createdSkill := db.DB.Create(&skill)
	err := createdSkill.Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "DB Error",
		})
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
	encodeErr := json.NewEncoder(c.Writer).Encode(&skill)
	if encodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": encodeErr,
		})
	}

}

func DeleteSkillHandler(c *gin.Context) {
	var skill models.Skill
	id := c.Param("id")
	db.DB.Find(&skill, id).Delete(&skill)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successful",
		"Value":   skill,
	})
}
