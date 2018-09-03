package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"strconv"
)

type IncomeController struct {
	gormDB gorm.DB
}

func NewIncomeController(db *gorm.DB) IncomeController {
	return IncomeController{gormDB: *db}
}

func (incomeController IncomeController) GetIncomeById(c *gin.Context) {
	var income models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	userId := int(jwt.ExtractClaims(c)["user_id"].(float64))
	if err != nil || userId == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	incomeController.gormDB.Where("id = ? AND user_id = ?", id, userId).First(&income, id)
	if income.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income": income})
}

func (incomeController IncomeController) AddIncome(c *gin.Context) {
	var income models.Income
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	c.BindJSON(&income)
	income.UserID = userId
	incomeController.gormDB.Create(&income)
	c.JSON(http.StatusCreated, gin.H{"income": income})
}

func (incomeController IncomeController) UpdateIncomeById(c *gin.Context) {
	var income models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	incomeController.gormDB.Where("id = ? AND user_id = ?", id, userId).First(&income, id)
	if income.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&income)
	incomeController.gormDB.Update(&income)
	c.JSON(http.StatusCreated, gin.H{"income": income})
}

func (incomeController IncomeController) DeleteIncomeById(c *gin.Context) {
	var income models.Income
	c.BindJSON(&income)
	incomeController.gormDB.Delete(&income)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}


func (incomeController IncomeController) UpdateIncomeByIdAndJWT(c *gin.Context) {
	var income models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	incomeController.gormDB.Where("id = ? AND user_id = ?", income.ID, userId).First(&income, id)
	if income.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&income)
	incomeController.gormDB.Update(&income)
	c.JSON(http.StatusCreated, gin.H{"income": income})
}

func (incomeController IncomeController) DeleteIncomeByIdAndJWT(c *gin.Context) {
	var income models.Income
	c.BindJSON(&income)
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	notFound := incomeController.gormDB.Where("id = ? AND user_id = ?", income.ID, userId).Delete(&income).RecordNotFound()
	if notFound {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
	}
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}