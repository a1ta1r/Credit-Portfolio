package models

import (
	"time"
)

type PaymentPlan struct {
	ID             uint      `json:"-"`
	Title          string    `json:"title"`
	LoanAmount     float64   `json:"loan_amount"`
	NumberOfMonths int       `json:"number_of_months"`
	InterestRate   float64   `json:"interest_rate"`
	StartDate      time.Time `json:"start_date"`
	MonthlyFee     Fee       `json:"monthly_fee"`
	PaymentsList   []Payment `json:"payments_list"`
}
