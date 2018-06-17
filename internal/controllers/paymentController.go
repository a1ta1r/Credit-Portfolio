package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type PaymentController struct {
	db gorm.DB
}

func NewPaymentController(db *gorm.DB) PaymentController {
	return PaymentController{db: *db}
}

func (pc PaymentController) GetPaymentsByPlan(c *gin.Context) {
	var payments []models.Payment
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	pc.db.Where("PaymentPlanID = ?", id).Find(&payments)
	c.JSON(http.StatusOK, payments)
}

func (pc PaymentController) GetPayment(c *gin.Context) {
	var payment models.Payment
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return
	}
	pc.db.First(&payment, id)
	if payment.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func (pc PaymentController) AddPayment(c *gin.Context) {
	var paymentPlan models.PaymentPlan
	c.BindJSON(&paymentPlan)
	pc.db.Create(&paymentPlan)
	c.JSON(http.StatusCreated, gin.H{"paymentPlan": paymentPlan})
}

func (pc PaymentController) DeletePayment(c *gin.Context) {
	var payment models.Payment
	c.BindJSON(&payment)
	if err := pc.db.Delete(&payment); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.NotFound})
	}
	c.JSON(http.StatusOK, gin.H{"message": utils.ResDeleted})
}
