package auth

import (
	"errors"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/roles"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/appleboy/gin-jwt.v2"
	"time"
)

var errInvalidCredentials = errors.New("username or password is invalid")

type JwtWrapper struct {
	userService services.UserService
}

func NewJwtWrapper(userService services.UserService) JwtWrapper {
	return JwtWrapper{userService: userService}
}

func (w JwtWrapper) GetJwtMiddleware(role roles.Role) jwt.GinJWTMiddleware {
	var authFunc = func(c *gin.Context) (interface{}, error) { return "", nil }
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
		Timeout:       time.Hour * 24 * 7,
		MaxRefresh:    time.Hour * 24,
		Authenticator: authFunc,
		PayloadFunc:   w.Payload,
	}
	return jwtMiddleware
}

func (w JwtWrapper) userRoleAuthFunc(c *gin.Context) (interface{}, error) {
	var users []entities.User
	var testUser entities.User
	c.BindJSON(&testUser)
	username := testUser.Username
	password := testUser.Password
	users = w.userService.GetUsers()
	var err error = nil
	for i := 0; i < len(users); i++ {
		err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && users[i].Role == roles.Basic && err == nil {
			return username, err
		}
	}
	return "", errInvalidCredentials
}

func (w JwtWrapper) adminRoleAuthFunc(c *gin.Context) (interface{}, error) {
	var users []entities.User
	users = w.userService.GetUsers()
	var err error = nil
	var testUser entities.User
	c.BindJSON(&testUser)
	username := testUser.Username
	password := testUser.Password
	for i := 0; i < len(users); i++ {
		err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && users[i].Role == roles.Admin && err == nil {
			return username, err
		}
	}
	return "", errInvalidCredentials
}

func (w JwtWrapper) merchantRoleAuthFunc(c *gin.Context) (interface{}, error) {
	var users []entities.User
	users = w.userService.GetUsers()
	var testUser entities.User
	c.BindJSON(&testUser)
	username := testUser.Username
	password := testUser.Password
	var err error = nil
	for i := 0; i < len(users); i++ {
		err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && users[i].Role == roles.Ads && err == nil {
			return username, err
		}
	}
	return "", errInvalidCredentials
}

func (w *JwtWrapper) Payload(username interface{}) jwt.MapClaims {
	var user = w.userService.GetUserByUsername(username.(string))
	return jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"user_id":  user.ID,
	}
}
