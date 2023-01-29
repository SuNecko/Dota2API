package main

import (
	"ApiCRUD/handlers"
	"github.com/gin-gonic/gin"
)

func setEndPoints(router *gin.Engine) {
	router.GET("/", handlers.HomeHandler)

	// Heroes Endpoints
	router.GET("/heroes", handlers.GetHeroesHandler)
	router.GET("/heroes/:name", handlers.GetHeroHandler)
	router.POST("/heroes", handlers.PostHeroHandler)
	router.DELETE("/heroes/:id", handlers.DeleteHeroHandler)

	//Items Endpoints
	router.GET("/items", handlers.GetItemsHandler)
	router.GET("/items/:name", handlers.GetItemHandler)
	router.POST("/items", handlers.PostItemHandler)
	router.DELETE("/items/:id", handlers.DeleteItemHandler)

	//Skills Endpoints
	router.GET("/skills", handlers.GetSkillsHandler)
	router.GET("/skills/:name", handlers.GetSkillHandler)
	router.POST("/skills", handlers.PostSkillsHandler)
	router.DELETE("/skills/:id", handlers.DeleteSkillHandler)
}
