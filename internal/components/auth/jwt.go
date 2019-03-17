package auth

import (
	"errors"
	"fmt"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/storages"
	le "github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	ls "github.com/a1ta1r/Credit-Portfolio/internal/components/loans/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/roles"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/appleboy/gin-jwt.v2"
	"time"
)

var errInvalidCredentials = errors.New("username or password is invalid")

type JwtWrapper struct {
	userService ls.UserService
	advStorage  storages.AdvertiserStorage
}

func NewJwtWrapper(userService ls.UserService, advStorage storages.AdvertiserStorage) JwtWrapper {
	return JwtWrapper{userService: userService, advStorage: advStorage}
}

func (w JwtWrapper) GetJwtMiddleware(role roles.Role) jwt.GinJWTMiddleware {
	jwtMiddleware := jwt.GinJWTMiddleware{
		Realm:         "robreid.io",
		Key:           []byte("portfolio-on-credit-very-very-very-secret-key"),
		Timeout:       time.Hour * 24 * 7,
		MaxRefresh:    time.Hour * 24,
		Authenticator: w.userRoleAuthFunc,
		PayloadFunc:   w.Payload,
	}
	return jwtMiddleware
}

func (w JwtWrapper) userRoleAuthFunc(c *gin.Context) (interface{}, error) {
	var users []le.User
	var testUser le.User
	c.BindJSON(&testUser)
	username := testUser.Username
	password := testUser.Password
	users = w.userService.GetUsers()
	var err error = nil
	for i := 0; i < len(users); i++ {
		err = bcrypt.CompareHashAndPassword([]byte(users[i].Password), []byte(password))
		if username == users[i].Username && users[i].Role == roles.Basic && err == nil {
			role := fmt.Sprint(users[i].Role)
			return map[string]interface{}{"id": users[i].ID, "username": username, "role": role}, err
		}
	}

	var advertisers []entities.Advertiser
	advertisers, _ = w.advStorage.GetAdvertisers()
	for i := 0; i < len(advertisers); i++ {
		err = bcrypt.CompareHashAndPassword([]byte(advertisers[i].Password), []byte(password))
		if username == advertisers[i].Username && advertisers[i].Role == roles.Ads && err == nil {
			role := fmt.Sprint(advertisers[i].Role)
			return map[string]interface{}{"id": advertisers[i].ID, "username": username, "role": role}, err
		}
	}

	return "", errInvalidCredentials
}

func (w JwtWrapper) adminRoleAuthFunc(c *gin.Context) (interface{}, error) {
	var users []le.User
	users = w.userService.GetUsers()
	var err error = nil
	var testUser le.User
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
	var users []le.User
	users = w.userService.GetUsers()
	var testUser le.User
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

	var advertisers []entities.Advertiser
	advertisers, _ = w.advStorage.GetAdvertisers()
	var testAdv entities.Advertiser
	c.BindJSON(&testAdv)
	username = testAdv.Username
	password = testAdv.Password
	for i := 0; i < len(advertisers); i++ {
		err = bcrypt.CompareHashAndPassword([]byte(advertisers[i].Password), []byte(password))
		if username == advertisers[i].Username && advertisers[i].Role == roles.Ads && err == nil {
			return username, err
		}
	}

	return "", errInvalidCredentials
}

func (w *JwtWrapper) Payload(user interface{}) jwt.MapClaims {
	id := (user.(map[string]interface{}))["id"]
	username := (user.(map[string]interface{}))["username"]
	role := (user.(map[string]interface{}))["role"]
	return jwt.MapClaims{
		"user_id":  id,
		"username": username,
		"role":     role,
	}
}
