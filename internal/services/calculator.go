package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"math"
)

type Calculator struct {
}

func (calculator Calculator) Calculate() {

	var a = Calculator{}
	var b = models.PaymentPlan{}
	b.NumberOfMonths = 12
	b.LoanAmount = 1000000
	b.InterestRate = 0.12

	var c = a.calculateDifferentiated(b)

	for _, x := range c  {
		println(x)
	}

}

func (calculator Calculator) calculateEqual(paymentPlan models.PaymentPlan) map[int]float64 {

	month := float64(paymentPlan.NumberOfMonths)
	percent := paymentPlan.InterestRate / 12

	koef := (percent * math.Pow(1+percent, month)) / (math.Pow(1+percent, month) - 1)

	sum := paymentPlan.LoanAmount * koef

	paymentPlan.ListPayments = make(map[int]float64)

	for i := 0; i < paymentPlan.NumberOfMonths; i++ {
		paymentPlan.ListPayments[i] = sum
	}
	return paymentPlan.ListPayments
}

func (calculator Calculator) calculateDifferentiated(paymentPlan models.PaymentPlan) map[int]float64 {

	baseFee := paymentPlan.LoanAmount / float64(paymentPlan.NumberOfMonths)

	paymentPlan.ListPayments = make(map[int]float64)

	for i := 0; i < paymentPlan.NumberOfMonths; i++ {
		paymentPlan.ListPayments[i] = baseFee + (paymentPlan.LoanAmount - baseFee*float64(i))*paymentPlan.InterestRate / float64(paymentPlan.NumberOfMonths)
	}

	return paymentPlan.ListPayments
}

