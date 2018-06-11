package models

import "time"

type Expense struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	User      User       `json:"-"`
	UserID    uint       `json:"user_id"`
	Reason    string     `json:"reason"`
	Amount    float64    `json:"amount"`
}
