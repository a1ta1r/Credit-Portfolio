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
	Role         Role          `json:"role"`
	RoleID       uint          `json:"roleId"`
	PaymentPlans []PaymentPlan `json:"paymentPlans"`
	Incomes      []Income      `json:"incomes"`
	Expenses     []Expense     `json:"expenses"`
}
