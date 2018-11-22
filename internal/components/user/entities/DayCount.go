package entities

import "time"

type DayCount struct {
	Date  time.Time `json:"date"`
	Count int       `json:"count"`
}
