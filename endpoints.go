package main

import (
	"ApiCRUD/handlers"
	"github.com/gin-gonic/gin"
)

func setEndPoints(router *gin.Engine) {
	router.GET("/", handlers.HomeHandler)

	group := router.Group("api/v1")

	// Heroes Endpoints
	group.GET("/heroes", handlers.GetHeroesHandler)
	group.GET("/heroes/:name", handlers.GetHeroHandler)
	group.GET("/heroes/attribute/:attribute", handlers.GetHeroByAttributeHandler)
	group.POST("/heroes", handlers.PostHeroHandler)
	group.DELETE("/heroes/:id", handlers.DeleteHeroHandler)

	//Items Endpoints
	group.GET("/items", handlers.GetItemsHandler)
	group.GET("/items/:name", handlers.GetItemHandler)
	group.POST("/items", handlers.PostItemHandler)
	group.DELETE("/items/:id", handlers.DeleteItemHandler)

	//Skills Endpoints
	group.GET("/skills", handlers.GetSkillsHandler)
	group.GET("/skills/:name", handlers.GetSkillHandler)
	group.POST("/skills", handlers.PostSkillsHandler)
	group.DELETE("/skills/:id", handlers.DeleteSkillHandler)
}
