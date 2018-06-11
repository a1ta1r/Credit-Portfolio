package app

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/handlers"
	"gopkg.in/appleboy/gin-jwt.v2"
	"github.com/a1ta1r/Credit-Portfolio/internal/controllers"
)

type JwtWrapper struct {
	userController controllers.UserController
}

func NewJwtWrapper(userController controllers.UserController) JwtWrapper {
	return JwtWrapper{userController: userController}
}

func (w JwtWrapper) GetJwtMiddleware() *jwt.GinJWTMiddleware {
	jwtMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "robreid.io",
		Key:           []byte("portfolio-on-credit-very-very-very-secret-key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: w.authenticatorFunc,
		PayloadFunc:   handlers.Payload,
	}
	return jwtMiddleware
}

func (w JwtWrapper) authenticatorFunc(username string, password string, c *gin.Context) (string, bool) {
	var users []models.User
	users, _, _ = w.userController.GetUsersArray(c)
	for i := 0; i < len(users); i++ {
		if username == users[i].Username && password == users[i].Password {
			return username, true
		}
	}
	return "", false
}
