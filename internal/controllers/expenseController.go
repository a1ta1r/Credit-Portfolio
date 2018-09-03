package controllers


import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type ExpenseController struct {
	gormDB gorm.DB
}

func NewExpenseController(db *gorm.DB) ExpenseController {
	return ExpenseController{gormDB: *db}
}

func (expenseController ExpenseController) GetExpenseById(c *gin.Context) {
	var expense models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	expenseController.gormDB.Preload("Payments").First(&expense, id)
	if expense.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expense": expense})
}

func (expenseController ExpenseController) AddExpense(c *gin.Context) {
	var expense models.Income
	c.BindJSON(&expense)
	expenseController.gormDB.Create(&expense)
	c.JSON(http.StatusCreated, gin.H{"expense": expense})
}

func (expenseController ExpenseController) UpdateExpenseById(c *gin.Context) {
	var expense models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	expenseController.gormDB.Preload("Payments").First(&expense, id)
	if expense.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&expense)
	expenseController.gormDB.Update(&expense)
	c.JSON(http.StatusCreated, gin.H{"expense": expense})
}

func (expenseController ExpenseController) DeleteExpenseById(c *gin.Context) {
	var expense models.Income
	c.BindJSON(&expense)
	expenseController.gormDB.Delete(&expense)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}
