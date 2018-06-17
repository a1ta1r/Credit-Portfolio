package models

import "time"

type TimePeriod struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name" gorm:"type:varchar(100);unique_index"`
}
