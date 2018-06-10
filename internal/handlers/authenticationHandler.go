package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	jwt "gopkg.in/appleboy/gin-jwt.v2"
)

func Hello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.String(http.StatusOK, "id: %s\nrole: %s", claims["id"], claims["role"])
}

func Authenticate(email string, password string, c *gin.Context) (string, bool) {
	if email == "credit-portfolio@gmail.com" && password == "credit-portfolio-password" {
		return email, true
	}
	return "", false
}

func Payload(email string) map[string]interface{} {
	return map[string]interface{}{
		"id":   "1349",
		"role": "admin",
	}
}
