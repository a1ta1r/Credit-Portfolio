package models

import "time"

type Payment struct {
	PaymentDate time.Time `json:"payment_date"`
	PaymentAmount  float64 `json:"payment_amount"`
}
