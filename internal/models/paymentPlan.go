package models

import (
	"time"
)

type PaymentPlan struct {
	ID            uint        `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
	DeletedAt     *time.Time  `sql:"index" json:"-"`
	Title         string      `json:"title"`
	User          User        `json:"-"`
	UserID        uint        `json:"user_id"`
	Bank          Bank        `json:"-"`
	BankID        uint        `json:"bank_id"`
	Currency      Currency    `json:"-"`
	CurrencyID    uint        `json:"currency_id"`
	PaymentType   PaymentType `json:"-"`
	PaymentTypeID uint        `json:"payment_type_id"`
	Amount        float64     `json:"payment_amount"`
	InterestRate  float64     `json:"interest_rate"`
	Months        uint        `json:"number_of_months"`
	StartDate     time.Time   `json:"start_date"`
	Payments      []Payment   `json:"payments"`
}
