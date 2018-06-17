package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/appleboy/gin-jwt.v2"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	gormDB gorm.DB
}

func NewUserController(pg *gorm.DB) UserController {
	return UserController{gormDB: *pg}
}

func (uc UserController) GetUserByName(c *gin.Context) {
	var user = uc.GetUserByUsername(c)
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return
	}
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

func (uc UserController) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		log.Println(err)
		panic(err)
	}
	var dbUser models.User
	uc.gormDB.
		Preload("Role").
		Preload("PaymentPlans").
		Preload("Expenses").
		Preload("Incomes").
		Preload("PaymentPlans.Payments").
		Table("users").Where("id = ?", user.ID).First(&dbUser)
	uc.gormDB.Model(&dbUser).Association("Incomes").Clear()
	uc.gormDB.Model(&dbUser).Association("Expenses").Clear()
	uc.gormDB.Model(&dbUser).Association("PaymentPlans").Clear()
	merged := mergeUsers(dbUser, user)
	uc.gormDB.Save(&merged)
	merged.Password = ""
	c.JSON(http.StatusOK, &merged)
}

//TODO нельзя создать юзера, чей мейл уже юзался, даже если юзер удалён!
func (uc UserController) AddUserAnonymous(c *gin.Context) {
	var user *models.User
	c.BindJSON(&user)
	user.RoleID = 1 //user
	var mailPassword = user.Password
	user.Password = getPasswordHash(user.Password)
	if err := uc.gormDB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": utils.ResourceExists})
		return
	} else {
		mail(user.Email, user.Username, mailPassword)
		c.JSON(http.StatusCreated, user)
	}
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

func (uc UserController) GetUserModelByGinContext(c *gin.Context) models.User {
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

func (uc UserController) GetUserByUsername(c *gin.Context) models.User {
	var user models.User
	username := c.Param("username")
	uc.gormDB.
		Preload("Role").
		Preload("PaymentPlans").
		Preload("Expenses").
		Preload("Incomes").
		Preload("PaymentPlans.Payments").
		Table("users").Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return models.User{}
	}
	user.Password = ""
	return user
}

func (uc UserController) GetUserModelByLogin(userId string) *models.User {
	var user = uc.gormDB.First(&models.User{}, "Username = ?", userId)
	if user.Value.(*models.User).ID == 0 {
		return nil
	}
	_ = user.Value
	return user.Value.(*models.User)
}

func getPasswordHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func mergeUsers(init models.User, new models.User) models.User {
	merged := init
	if new.Password != "" {
		merged.Password = getPasswordHash(new.Password)
	}
	if new.Email != "" {
		merged.Email = new.Email
	}
	if new.RoleID != 0 {
		merged.RoleID = new.RoleID
	}
	if len(new.Expenses) > 0 {
		merged.Expenses = new.Expenses
	}
	if len(new.Incomes) > 0 {
		merged.Incomes = new.Incomes
	}
	if len(new.PaymentPlans) > 0 {
		merged.PaymentPlans = new.PaymentPlans
	}
	return merged
}

func (uc UserController) GetUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	var user = uc.GetUserModelByLogin(claims["id"].(string))
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
