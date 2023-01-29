package main

import (
	"ApiCRUD/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db.SetupDB()

	router := gin.Default()
	setEndPoints(router)
	log.Fatal(router.Run(":8080"))
}
