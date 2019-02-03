package entities

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/roles"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           uint          `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	DeletedAt    *time.Time    `json:"-"`
	Username     string        `json:"username" gorm:"type:varchar(100);unique_index"`
	Email        string        `json:"email" gorm:"type:varchar(100);unique_index"`
	Password     string        `json:"password,omitempty"`
	Role         roles.Role    `json:"role"`
	PaymentPlans []PaymentPlan `json:"paymentPlans",default:"[]"`
	Incomes      []Income      `json:"incomes",default:"[]"`
	Expenses     []Expense     `json:"expenses",default:"[]"`
	LastSeen     *time.Time    `json:"lastSeen"`
}

func (u User) GetHashedPassword() string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
