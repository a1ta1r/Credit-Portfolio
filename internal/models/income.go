package models

import "time"

type Income struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
	User          User       `json:"-"`
	UserID        uint       `json:"userId"`
	Reason        string     `json:"reason"`
	Amount        float64    `json:"amount"`
	StartDate     time.Time  `json:"startDate"`
	IsRepeatable  bool       `json:"isRepeatable"`
	Frequency     uint       `json:"frequency"`
	PaymentPeriod TimePeriod `json:"paymentPeriod"`
}
