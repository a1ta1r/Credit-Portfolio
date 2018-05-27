package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"math"
	"time"
)

type Calculator struct {
}

func (calculator Calculator) Calculate() models.PaymentPlan {
	var b = models.PaymentPlan{
		LoanAmount:     1000000,
		NumberOfMonths: 12,
		InterestRate:   0.12,
		StartDate:      time.Now(),
		ID:             1,
		Title:          "Test Payment Plan",
	}
	return calculator.calculateDifferentiated(b)
}

func (calculator Calculator) calculateEqual(paymentPlan models.PaymentPlan) models.PaymentPlan {

	month := float64(paymentPlan.NumberOfMonths)
	percent := paymentPlan.InterestRate / 12

	koef := (percent * math.Pow(1+percent, month)) / (math.Pow(1+percent, month) - 1)

	sum := paymentPlan.LoanAmount * koef

	paymentPlan.PaymentsList = []models.Payment{}

	for i := 0; i < paymentPlan.NumberOfMonths; i++ {
		paymentDate := paymentPlan.StartDate.AddDate(0, i, 0)
		payment := models.Payment{
			PaymentDate:   paymentDate,
			PaymentAmount: sum,
		}
		paymentPlan.PaymentsList = append(paymentPlan.PaymentsList, payment)
	}
	return paymentPlan
}

func (calculator Calculator) calculateDifferentiated(paymentPlan models.PaymentPlan) models.PaymentPlan {

	baseFee := paymentPlan.LoanAmount / float64(paymentPlan.NumberOfMonths)

	paymentPlan.PaymentsList = []models.Payment{}

	for i := 0; i < paymentPlan.NumberOfMonths; i++ {
		sum := baseFee + (paymentPlan.LoanAmount-baseFee*float64(i))*paymentPlan.InterestRate/float64(paymentPlan.NumberOfMonths)
		paymentDate := paymentPlan.StartDate.AddDate(0, i, 0)
		payment := models.Payment{
			PaymentDate:   paymentDate,
			PaymentAmount: sum,
		}
		paymentPlan.PaymentsList = append(paymentPlan.PaymentsList, payment)
	}

	return paymentPlan
}
