package models

import "time"

type User struct {
	ID           uint          `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	DeletedAt    *time.Time    `sql:"index" json:"-"`
	Username     string        `json:"username" gorm:"type:varchar(100);unique_index"`
	Email        string        `json:"email" gorm:"type:varchar(100);unique_index"`
	Password     string        `json:"password,omitempty"`
	Role         Role          `json:"-"`
	RoleID       uint          `json:"role_id"`
	PaymentPlans []PaymentPlan `json:"payment_plans"`
}
