package auth

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/roles"
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

func (w JwtWrapper) GetJwtMiddleware(role roles.Role) jwt.GinJWTMiddleware {
	var authFunc = func(username string, password string, c *gin.Context) (string, bool) { return "", false }
	switch role {
	case roles.Basic:
		authFunc = w.userRoleAuthFunc
		break
	case roles.Admin:
		authFunc = w.adminRoleAuthFunc
		break
	case roles.Ads:
		authFunc = w.merchantRoleAuthFunc
		break
	}
	jwtMiddleware := jwt.GinJWTMiddleware{
		Realm:         "robreid.io",
		Key:           []byte("portfolio-on-credit-very-very-very-secret-key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: authFunc,
		PayloadFunc:   w.Payload,
	}
	return jwtMiddleware
}

func (w JwtWrapper) userRoleAuthFunc(username string, password string, c *gin.Context) (string, bool) {
	var users []entities.User
	users = w.userService.GetUsers()
	for i := 0; i < len(users); i++ {
		var err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && err == nil {
			return username, true
		}
	}
	return "", false
}

func (w JwtWrapper) adminRoleAuthFunc(username string, password string, c *gin.Context) (string, bool) {
	var users []entities.User
	users = w.userService.GetUsers()
	for i := 0; i < len(users); i++ {
		var err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && users[i].Role == roles.Ads && err == nil {
			return username, true
		}
	}
	return "", false
}

func (w JwtWrapper) merchantRoleAuthFunc(username string, password string, c *gin.Context) (string, bool) {
	var users []entities.User
	users = w.userService.GetUsers()
	for i := 0; i < len(users); i++ {
		var err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && users[i].Role == roles.Ads && err == nil {
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
		"user_id":  user.ID,
	}
}
