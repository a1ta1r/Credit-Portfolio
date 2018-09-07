package models

import (
	"time"
)

type PaymentPlan struct {
	ID            uint        `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
	Title         string      `json:"title"`
	User          User        `json:"-"`
	UserID        uint        `json:"userId"`
	Bank          Bank        `json:"-"`
	BankID        uint        `json:"bankId"`
	Currency      Currency    `json:"-"`
	CurrencyID    uint        `json:"currencyId"`
	PaymentType   PaymentType `json:"paymentType"`
	Amount        float64     `json:"paymentAmount"`
	InterestRate  float64     `json:"interestRate"`
	Months        uint        `json:"numberOfMonths"`
	StartDate     time.Time   `json:"startDate"`
	Payments      []Payment   `json:"paymentList"`
	TotalPaymentAmount float64 `json:"totalPaymentAmount"`
}
