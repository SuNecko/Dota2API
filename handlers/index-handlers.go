package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to my APIRest")
}
