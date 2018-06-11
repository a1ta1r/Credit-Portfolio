package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/appleboy/gin-jwt.v2"
)

func Hello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(http.StatusOK, gin.H{"id": claims["id"], "role": claims["role"]})
}

func Payload(email string) map[string]interface{} {
	return map[string]interface{}{
		"id":   "1349",
		"role": "admin",
	}
}
