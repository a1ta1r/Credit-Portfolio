package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type UserController struct {
	gormDB gorm.DB
}

func (uc UserController) GetUsers(c *gin.Context) {
	var users []models.User
	uc.gormDB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc UserController) GetUser(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return
	}
	uc.gormDB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

//TODO правильно создавать связи в с role в бд
func (uc UserController) AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	uc.gormDB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (uc UserController) DeleteUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	uc.gormDB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": utils.ResDeleted})
}

func NewUserController(pg *gorm.DB) UserController {
	return UserController{gormDB: *pg}
}
