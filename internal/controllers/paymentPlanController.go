package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type PaymentPlanController struct {
	gormDB gorm.DB
}

func NewPaymentPlanController(db *gorm.DB) PaymentPlanController {
	return PaymentPlanController{gormDB: *db}
}

func (paymentPlanController PaymentPlanController) GetPaymentPlans(c *gin.Context) {
	var paymentPlans []models.PaymentPlan
	limit, offset := int64(-1), int64(0)
	reqLimit, _ := strconv.ParseInt(c.Query("limit"), 10, 32)
	reqOffset, _ := strconv.ParseInt(c.Query("offset"), 10, 32)
	if reqLimit > 0 {
		limit = reqLimit
	}
	if reqOffset > 0 {
		offset = reqOffset
	}
	paymentPlanController.gormDB.Offset(offset).Limit(limit).Find(&paymentPlans)
	c.JSON(http.StatusOK, gin.H{
		"limit":        limit,
		"offset":       offset,
		"paymentPlans": paymentPlans,
	})
}

func (paymentPlanController PaymentPlanController) GetPaymentPlan(c *gin.Context) {
	var paymentPlan models.PaymentPlan
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return
	}
	paymentPlanController.gormDB.Preload("Payments").First(&paymentPlan, id)
	if paymentPlan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"paymentPlan": paymentPlan})
}

func (paymentPlanController PaymentPlanController) AddPaymentPlan(c *gin.Context) {
	var paymentPlan models.PaymentPlan
	c.BindJSON(&paymentPlan)
	paymentPlanController.gormDB.Create(&paymentPlan)
	c.JSON(http.StatusCreated, gin.H{"paymentPlan": paymentPlan})
}

func (paymentPlanController PaymentPlanController) DeletePaymentPlan(c *gin.Context) {
	var paymentPlan models.PaymentPlan
	c.BindJSON(&paymentPlan)
	paymentPlanController.gormDB.Delete(&paymentPlan)
	c.JSON(http.StatusOK, gin.H{"message": utils.ResDeleted})
}
