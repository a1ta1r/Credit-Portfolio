package models

import (
	"time"
)

type PaymentPlan struct {
	LoanAmount     float64
	NumberOfMonths int
	InterestRate   float64
	StartDate      time.Time
	MonthlyFee     Fee
	ListPayments   map[int]float64

}
