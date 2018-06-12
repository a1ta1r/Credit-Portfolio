package models

import (
	"time"
)

type Payment struct {
	ID            uint        `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
	DeletedAt     *time.Time  `sql:"index" json:"-"`
	PaymentPlan   PaymentPlan `json:"-"`
	PaymentPlanID uint        `json:"paymentPlanId"`
	PaymentDate   time.Time   `json:"paymentDate"`
	PaymentAmount float64     `json:"paymentAmount"`
}
