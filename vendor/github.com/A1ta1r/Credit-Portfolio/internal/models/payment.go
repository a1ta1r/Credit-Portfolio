package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	PaymentPlan   PaymentPlan `json:"-"`
	PaymentPlanID uint        `json:"payment_plan_id"`
	PaymentDate   time.Time   `json:"payment_date"`
	PaymentAmount float64     `json:"payment_amount"`
}
