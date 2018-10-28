package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"strconv"
)

type PaymentPlanController struct {
	gormDB      gorm.DB
	userService services.UserService
}

func NewPaymentPlanController(db *gorm.DB, userService services.UserService) PaymentPlanController {
	return PaymentPlanController{*db, userService}
}

func (paymentPlanController PaymentPlanController) GetPaymentPlans(c *gin.Context) {
	var paymentPlans []entities.PaymentPlan
	userId := int(jwt.ExtractClaims(c)["user_id"].(float64))
	paymentPlanController.gormDB.Where("user_id = ?", userId).Find(&paymentPlans)
	c.JSON(http.StatusOK, gin.H{
		"paymentPlans": paymentPlans,
	})
}

func (paymentPlanController PaymentPlanController) GetPaymentPlan(c *gin.Context) {
	var paymentPlan entities.PaymentPlan
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	userId := int(jwt.ExtractClaims(c)["user_id"].(float64))
	if err != nil || userId == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	paymentPlanController.gormDB.Preload("Payments").Where("id = ? AND user_id = ?", id, userId).First(&paymentPlan, id)
	if paymentPlan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"paymentPlan": paymentPlan})
}

func (paymentPlanController PaymentPlanController) AddPaymentPlan(c *gin.Context) {
	var paymentPlan entities.PaymentPlan
	c.BindJSON(&paymentPlan)
	paymentPlan.UserID = uint(jwt.ExtractClaims(c)["user_id"].(float64))
	paymentPlanController.gormDB.Create(&paymentPlan)
	c.JSON(http.StatusCreated, gin.H{"paymentPlan": paymentPlan})
}

func (paymentPlanController PaymentPlanController) UpdatePaymentPlan(c *gin.Context) {
	var paymentPlan entities.PaymentPlan
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	userId := int(jwt.ExtractClaims(c)["user_id"].(float64))
	if err != nil || userId == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	paymentPlanController.gormDB.Preload("Payments").Where("id = ? AND user_id = ?", id, userId).First(&paymentPlan, id)
	if paymentPlan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&paymentPlan)
	paymentPlanController.gormDB.Create(&paymentPlan)
	c.JSON(http.StatusCreated, gin.H{"paymentPlan": paymentPlan})
}

func (paymentPlanController PaymentPlanController) DeletePaymentPlan(c *gin.Context) {
	var paymentPlan entities.PaymentPlan
	c.BindJSON(&paymentPlan)
	paymentPlanController.gormDB.Delete(&paymentPlan)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}
