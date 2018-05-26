package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//checks if services are still alive and returns 200
//returns 503 otherwise
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
