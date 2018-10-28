package entities

import "time"

type Currency struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name" gorm:"type:varchar(100);unique_index"`
	Symbol    string    `json:"symbol"`
}
