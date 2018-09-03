package models

import "time"

type Income struct {
	ID             uint        `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time   `json:"-"`
	UpdatedAt      time.Time   `json:"-"`
	User           User        `json:"user"`
	UserID         uint        `json:"userId"`
	Reason         string      `json:"reason"`
	Amount         float64     `json:"amount"`
	StartDate      time.Time   `json:"startDate"`
	IsRepeatable   bool        `json:"isRepeatable"` //рекуррентный платеж или нет
	Frequency      *uint       `json:"frequency"`
	PaymentPeriod  *TimePeriod `json:"paymentPeriod"`
	RecurrentCount *uint       `json:"recurrentCount"` //число повторений в TimePeriod(4 недели, 12 лет)
}

func (i Income) transform() AgendaElement {
	return AgendaElement{}
}
