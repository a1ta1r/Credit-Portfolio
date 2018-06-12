package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserController struct {
	gormDB gorm.DB
}

func NewUserController(pg *gorm.DB) UserController {
	return UserController{gormDB: *pg}
}

func (uc UserController) GetUser(c *gin.Context) {
	var user = uc.GetUserEntityByGinContext(c)
	c.JSON(http.StatusOK, user)
}

func (uc UserController) GetUsers(c *gin.Context) {
	var users []models.User
	users, _, _ = uc.GetUsersArray(c)
	for i := 0; i < len(users); i++ {
		users[i].Password = ""
	}
	c.JSON(http.StatusOK, users)
}

func (uc UserController) AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	uc.gormDB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (uc UserController) AddUserAnonymous(c *gin.Context) {
	var user *models.User
	c.BindJSON(&user)
	user.RoleID = 1 //user
	user.Password = getPasswordHash(user.Password)
	uc.gormDB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (uc UserController) DeleteUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return
	}
	uc.gormDB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{"message": utils.ResDeleted})
}

func (uc UserController) GetUsersArray(c *gin.Context) ([]models.User, int64, int64) {
	var users []models.User
	limit, offset := int64(-1), int64(0)
	reqLimit, _ := strconv.ParseInt(c.Query("limit"), 10, 32)
	reqOffset, _ := strconv.ParseInt(c.Query("offset"), 10, 32)
	if reqLimit > 0 {
		limit = reqLimit
	}
	if reqOffset > 0 {
		offset = reqOffset
	}
	uc.gormDB.Offset(offset).Limit(limit).Find(&users)
	return users, limit, offset
}

func (uc UserController) GetUserEntityByGinContext(c *gin.Context) models.User {
	var user models.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return models.User{}
	}
	uc.gormDB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return models.User{}
	}
	user.Password = ""
	return user
}

func (uc UserController) GetUserById(userId string) models.User {
	var user models.User
	id, _ := strconv.ParseUint(userId, 10, 32)
	uc.gormDB.First(&user, id)
	if user.ID == 0 {
		return models.User{}
	}
	return user
}

func getPasswordHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
