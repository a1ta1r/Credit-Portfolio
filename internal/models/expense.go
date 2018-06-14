package models

import "time"

type Expense struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	User      User       `json:"-"`
	UserID    uint       `json:"userId"`
	Reason    string     `json:"reason"`
	Amount    float64    `json:"amount"`
	EndDate   time.Time  `json:"endDate"`
}
