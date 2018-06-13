package app

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"gopkg.in/appleboy/gin-jwt.v2"
	"github.com/a1ta1r/Credit-Portfolio/internal/controllers"
	"golang.org/x/crypto/bcrypt"
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
		PayloadFunc:   w.Payload,
	}
	return jwtMiddleware
}

func (w JwtWrapper) authenticatorFunc(username string, password string, c *gin.Context) (string, bool) {
	var users []models.User
	users, _, _ = w.userController.GetUsersArray(c)
	for i := 0; i < len(users); i++ {
		var err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && err == nil {
			return username, true
		}
	}
	return "", false
}

func (w *JwtWrapper) Payload(userId string) map[string]interface{} {
	var user = w.userController.GetUserModelById(userId)
	return map[string]interface{}{
		"username": user.Username,
		"role": user.Role.Name,
	}
}
