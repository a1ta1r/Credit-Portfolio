package controllers


import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"strconv"
)

type IncomeController struct {
	gormDB gorm.DB
	userService services.UserService
}

func NewIncomeController(db *gorm.DB, userService services.UserService) IncomeController {
	return IncomeController{gormDB: *db, userService: userService}
}

func (incomeController IncomeController) GetIncomeById(c *gin.Context) {
	var income models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	incomeController.gormDB.Preload("Payments").First(&income, id)
	if income.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income": income})
}

//func (incomeController IncomeController) AddIncome(c *gin.Context) {
//	var income models.Income
//	c.BindJSON(&income)
//	incomeController.gormDB.Create(&income)
//	c.JSON(http.StatusCreated, gin.H{"income": income})
//}

func (incomeController IncomeController) UpdateIncomeById(c *gin.Context) {
	var income models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	incomeController.gormDB.Preload("Payments").First(&income, id)
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

func (incomeController IncomeController) AddIncome(c *gin.Context) {
	var income models.Income
	c.BindJSON(&income)
	income.UserID = uint(jwt.ExtractClaims(c)["user_id"].(float64))
	incomeController.gormDB.Create(&income)
	userId := int(jwt.ExtractClaims(c)["user_id"].(float64))
	incomeController.gormDB.Where("user_id = ?", userId).Find(&income)
	c.JSON(http.StatusCreated, gin.H{"income": income})
}