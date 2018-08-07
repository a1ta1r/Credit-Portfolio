package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           uint          `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	Username     string        `json:"username" gorm:"type:varchar(100);unique_index"`
	Email        string        `json:"email" gorm:"type:varchar(100);unique_index"`
	Password     string        `json:"password,omitempty"`
	Role         Role          `json:"role"`
	PaymentPlans []PaymentPlan `json:"paymentPlans"`
	Incomes      []Income      `json:"incomes"`
	Expenses     []Expense     `json:"expenses"`
}

func (u User) GetHashedPassword() string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
