package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"strconv"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) UserController {
	return UserController{userService: service}
}

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

func (uc UserController) GetUsers(c *gin.Context) {
	limit, offset := int64(-1), int64(0)
	reqLimit, _ := strconv.ParseInt(c.Query("limit"), 10, 32)
	reqOffset, _ := strconv.ParseInt(c.Query("offset"), 10, 32)
	if reqLimit > 0 {
		limit = reqLimit
	}
	if reqOffset > 0 {
		offset = reqOffset
	}
	users := uc.userService.GetUsers(offset, limit)
	for i := 0; i < len(users); i++ {
		users[i].Password = ""
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"count":  len(users),
		"users":  users,
	})
}

func (uc UserController) UpdateUser(c *gin.Context) {
	var user models.User
	c.ShouldBindWith(&user, binding.JSON)
	user = uc.userService.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"user":   user,
	})
}

func (uc UserController) AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	user.Role = models.Basic
	var mailPassword = user.Password
	user.Password = user.GetHashedPassword()
	user = uc.userService.CreateUser(user)
	mail(user.Email, user.Username, mailPassword)
	c.JSON(http.StatusCreated, user)
}

func (uc UserController) DeleteUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	uc.userService.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

func (uc UserController) GetUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id, _ := strconv.ParseInt(claims["id"].(string), 10, 32)
	user := uc.userService.GetUserByID(uint(id))
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
