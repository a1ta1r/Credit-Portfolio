package entities

import (
	"time"
)

type Payment struct {
	ID            uint        `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
	PaymentPlan   PaymentPlan `json:"-"`
	PaymentPlanID uint        `json:"paymentPlanId"`
	PaymentDate   time.Time   `json:"paymentDate"`
	PaymentAmount float64     `json:"paymentAmount"`
}

func (p Payment) Transform() AgendaElement {
	return AgendaElement{
		"Payment",
		p.ID,
		p.PaymentPlan.UserID,
		p.PaymentPlan.Title,
		p.PaymentAmount,
		p.PaymentDate,
	}
}

func (p Payment) TransformWithPeriod(from time.Time, to time.Time) []AgendaElement {
	if p.PaymentDate.After(from) && p.PaymentDate.Before(to) {
		return []AgendaElement{p.Transform()}
	} else {
		return []AgendaElement{}
	}
}
