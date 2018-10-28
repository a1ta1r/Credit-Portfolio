package entities

import "time"

type CreditCalculation struct {
	Title          string    `json:"title"`
	InterestRate   uint      `json:"interestRate"`
	NumberOfMonths uint      `json:"numberOfMonths"`
	PaymentAmount  uint      `json:"paymentAmount"`
	CreditType     string    `json:"creditType"`
	StartDate      time.Time `json:"startDate"`
}
