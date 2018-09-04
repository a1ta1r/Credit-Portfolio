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

type ExpenseController struct {
	gormDB gorm.DB
}

func NewExpenseController(db *gorm.DB) ExpenseController {
	return ExpenseController{gormDB: *db}
}

func (expenseController ExpenseController) GetExpensesByJWT(c *gin.Context) {
	var expenses []models.Expense
	userId := int(jwt.ExtractClaims(c)["user_id"].(float64))
	if userId == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	expenseController.gormDB.Where("user_id = ?", userId).Find(&expenses)
	if len(expenses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}

func (expenseController ExpenseController) GetExpenseById(c *gin.Context) {
	var expense models.Expense
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	userId := int(jwt.ExtractClaims(c)["user_id"].(float64))
	if err != nil || userId == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	expenseController.gormDB.Where("id = ? AND user_id = ?", id, userId).First(&expense, id)
	if expense.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expense": expense})
}

func (expenseController ExpenseController) AddExpense(c *gin.Context) {
	var expense models.Expense
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	c.BindJSON(&expense)
	expense.UserID = userId
	expenseController.gormDB.Create(&expense)
	c.JSON(http.StatusCreated, gin.H{"expense": expense})
}

func (expenseController ExpenseController) UpdateExpenseById(c *gin.Context) {
	var expense models.Expense
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	expenseController.gormDB.Where("id = ? AND user_id = ?", id, userId).First(&expense, id)
	if expense.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&expense)
	expenseController.gormDB.Update(&expense)
	c.JSON(http.StatusCreated, gin.H{"expense": expense})
}

func (expenseController ExpenseController) DeleteExpenseById(c *gin.Context) {
	var expense models.Expense
	c.BindJSON(&expense)
	expenseController.gormDB.Delete(&expense)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

func (expenseController ExpenseController) UpdateExpenseByIdAndJWT(c *gin.Context) {
	var expense models.Expense
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	expenseController.gormDB.Where("id = ? AND user_id = ?", expense.ID, userId).First(&expense, id)
	if expense.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&expense)
	expenseController.gormDB.Update(&expense)
	c.JSON(http.StatusCreated, gin.H{"expense": expense})
}

func (expenseController ExpenseController) DeleteExpenseByIdAndJWT(c *gin.Context) {
	var expense models.Expense
	c.BindJSON(&expense)
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	notFound := expenseController.gormDB.Where("id = ? AND user_id = ?", expense.ID, userId).Delete(&expense).RecordNotFound()
	if notFound {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
	}
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}