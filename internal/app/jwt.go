package app

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/appleboy/gin-jwt.v2"
	"time"
)

type JwtWrapper struct {
	userService services.UserService
}

func NewJwtWrapper(userService services.UserService) JwtWrapper {
	return JwtWrapper{userService: userService}
}

func (w JwtWrapper) GetJwtMiddleware() *jwt.GinJWTMiddleware {
	jwtMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "robreid.io",
		Key:           []byte("portfolio-on-credit-very-very-very-secret-key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: w.authenticatorFunc,
		PayloadFunc:   w.Payload,
	}
	return jwtMiddleware
}

func (w JwtWrapper) authenticatorFunc(username string, password string, c *gin.Context) (string, bool) {
	var users []models.User
	users = w.userService.GetUsers(0, -1)
	for i := 0; i < len(users); i++ {
		var err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && err == nil {
			return username, true
		}
	}
	return "", false
}

func (w *JwtWrapper) Payload(username string) map[string]interface{} {
	var user = w.userService.GetUserByUsername(username)
	return map[string]interface{}{
		"username": user.Username,
		"role":     user.Role,
	}
}
