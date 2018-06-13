package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type PaymentPlan struct {
	gorm.Model
	User          User
	UserID        uint
	Bank          Bank
	BankID        uint
	Currency      Currency
	CurrencyID    uint
	PaymentType   PaymentType
	PaymentTypeID uint
	Amount        float64
	InterestRate  float64
	Months        uint
	StartDate     time.Time
	Payments      []Payment
}
