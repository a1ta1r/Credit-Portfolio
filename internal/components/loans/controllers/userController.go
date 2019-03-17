package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/mailer"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/roles"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"

	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/services"
	"github.com/gin-gonic/gin"
	"gopkg.in/appleboy/gin-jwt.v2"
)

//UserController processes user-related HTTP requests
type UserController struct {
	userService services.UserService
}

//NewUserController returns an instance of UserController
func NewUserController(service services.UserService) UserController {
	return UserController{userService: service}
}

//GetUserByUsername returns a user entity associated with given username
func (uc UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user := uc.userService.GetUserByUsername(username)
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"user":   user,
	})
}

//GetUsers returns all users present in the database
func (uc UserController) GetUsers(c *gin.Context) {
	users := uc.userService.GetUsers()
	for i := 0; i < len(users); i++ {
		users[i].Password = ""
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"count":  len(users),
		"users":  users,
	})
}

//UpdateAdvertiser updates user data in database and returns new user entity in JSON
func (uc UserController) UpdateUser(c *gin.Context) {
	var user entities.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	user = uc.userService.GetUserByID(uint(id))
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}

	c.ShouldBindWith(&user, binding.JSON)
	user = uc.userService.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"user":   user,
	})
}

func (uc UserController) UpdateUserByJWT(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := uint(claims["user_id"].(float64))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	user := uc.userService.GetUserByID(id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}

	role := user.Role
	c.ShouldBindWith(&user, binding.JSON)
	user.Role = role
	user = uc.userService.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"user":   user,
	})
}

func (uc UserController) GetUser(c *gin.Context) {
	var user entities.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	user = uc.userService.GetUserByID(uint(id))
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

//AddUser creates new user entity and adds it to database
func (uc UserController) AddUser(c *gin.Context) {
	var user entities.User
	c.BindJSON(&user)
	user.Role = roles.Basic
	var mailPassword = user.Password
	user.Password = user.GetHashedPassword()
	user.LastSeen = time.Now()
	user, isOk := uc.userService.CreateUser(user)
	if !isOk {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": codes.ResourceExists})
		return
	}
	mailer.SendMail(user.Email, user.Username, mailPassword)
	user = uc.userService.GetUserByUsername(user.Username)
	c.JSON(http.StatusCreated, user)
}

//DeleteUser removes user by ID
func (uc UserController) DeleteUser(c *gin.Context) {
	var user entities.User
	c.BindJSON(&user)
	uc.userService.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

//GetUserByJWT returns JSON with currently authenticated user using JWT
func (uc UserController) GetUserByJWT(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["user_id"].(float64))
	user := uc.userService.GetUserByID(uint(id))
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
