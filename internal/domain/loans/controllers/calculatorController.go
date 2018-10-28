package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"gopkg.in/appleboy/gin-jwt.v2"
	"math"
	"net/http"
)

type CalculatorController struct {
	db gorm.DB
}

func NewCalculatorController(pg *gorm.DB) CalculatorController {
	return CalculatorController{db: *pg}
}

func (cc CalculatorController) CalculateCredit(c *gin.Context) {
	userId := uint(jwt.ExtractClaims(c)["user_id"].(float64))
	var paymentPlan models.PaymentPlan
	c.ShouldBindWith(&paymentPlan, binding.JSON)
	paymentPlan.UserID = userId
	if paymentPlan.PaymentType == models.Even {
		paymentPlan = CalculateCreditWithEqualPayments(paymentPlan)
	} else {
		paymentPlan = CalculateCreditWithDifferentiatedPayments(paymentPlan)
	}
	if cc.db.Create(&paymentPlan).Error != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.Unhealthy})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"paymentPlan": paymentPlan,
	})
}

func CalculateCreditWithEqualPayments(paymentPlan models.PaymentPlan) models.PaymentPlan {
	var percent = paymentPlan.InterestRate / 12 / 100
	var coefficient = (percent * math.Pow(1+percent, float64(paymentPlan.Months))) / (math.Pow(1+percent, float64(paymentPlan.Months)) - 1)
	var sum = paymentPlan.Amount * coefficient

	paymentPlan.Payments = []models.Payment{}
	currentMonth := paymentPlan.StartDate

	for i := 0; i < int(paymentPlan.Months); i++ {
		var payment = models.Payment{PaymentDate: currentMonth, PaymentAmount: sum}
		paymentPlan.Payments = append(paymentPlan.Payments, payment)
		currentMonth = currentMonth.AddDate(0, 1, 0)
	}
	paymentPlan.TotalPaymentAmount = sum * float64(paymentPlan.Months)
	return paymentPlan
}

func CalculateCreditWithDifferentiatedPayments(paymentPlan models.PaymentPlan) models.PaymentPlan {
	var baseFee = paymentPlan.Amount / float64(paymentPlan.Months)

	paymentPlan.Payments = []models.Payment{}
	paymentPlan.TotalPaymentAmount = 0
	currentMonth := paymentPlan.StartDate

	for i := 0; i < int(paymentPlan.Months); i++ {
		currentMonth = currentMonth.AddDate(0, 1, 0)
		var sum = baseFee + (paymentPlan.Amount-baseFee*float64(i))*paymentPlan.InterestRate/100/12
		var payment = models.Payment{PaymentDate: currentMonth, PaymentAmount: sum}
		paymentPlan.Payments = append(paymentPlan.Payments, payment)
		paymentPlan.TotalPaymentAmount += sum
	}
	return paymentPlan
}
