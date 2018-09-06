package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"time"
)

type CalculatorController struct {
	db gorm.DB
}

func NewCalculatorController(pg *gorm.DB) CalculatorController {
	return CalculatorController{db: *pg}
}

func (cc CalculatorController) CalculateCredit(c *gin.Context) {
	var creditCalculation models.CreditCalculation
	c.BindJSON(&creditCalculation)
	var paymentPlan models.PaymentPlan
	if creditCalculation.CreditType == "dafwd" {
		paymentPlan = CalculateCreditWithEqualPayments(float64(creditCalculation.InterestRate), float64(creditCalculation.NumberOfMonths), float64(creditCalculation.PaymentAmount), creditCalculation.StartDate)
	} else {
		paymentPlan = CalculateCreditWithDifferentiatedPayments(float64(creditCalculation.InterestRate), float64(creditCalculation.NumberOfMonths), float64(creditCalculation.PaymentAmount), creditCalculation.StartDate)
	}
	cc.db.Create(&paymentPlan)
}

func CalculateCreditWithEqualPayments(interestRate float64, numberOfMonths float64, paymentAmount float64, startDate time.Time) models.PaymentPlan {
	var percent = interestRate / 12 / 100
	var coefficient = (percent * math.Pow(1 + percent, float64(numberOfMonths))) / (math.Pow(1 + percent,  float64(numberOfMonths)) - 1)
	var sum = paymentAmount * coefficient
	var paymentPlan models.PaymentPlan
	paymentPlan.Payments = [] models.Payment{}
	for i := 0; i < int(numberOfMonths); i++ {
		var currentMonth = startDate.AddDate(0,1,0)
		var payment = models.Payment{PaymentDate: currentMonth, PaymentAmount: sum}
		paymentPlan.Payments = append(paymentPlan.Payments, payment)
	}
	return paymentPlan
}

func CalculateCreditWithDifferentiatedPayments(interestRate float64, numberOfMonths float64, paymentAmount float64, startDate time.Time) models.PaymentPlan {
	var baseFee = paymentAmount / numberOfMonths
	var paymentPlan models.PaymentPlan
	paymentPlan.Payments = [] models.Payment{}
	paymentPlan.Amount = 0

	for i := 0; i < int(numberOfMonths); i++ {
		var currentMonth = startDate.AddDate(0,1,0)
		var sum = baseFee + (paymentAmount - baseFee * float64(i)) * interestRate / 100 / 12
		var payment = models.Payment{PaymentDate: currentMonth, PaymentAmount: sum}
		paymentPlan.Payments = append(paymentPlan.Payments, payment)
	}
	return paymentPlan
}