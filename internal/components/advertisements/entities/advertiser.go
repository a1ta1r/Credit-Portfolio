package entities

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/roles"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Advertiser struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"createdAt" example:"2018-10-30T19:43:15.251038Z"`
	UpdatedAt   time.Time  `json:"updatedAt" example:"2018-10-30T19:43:15.251038Z"`
	Username    string     `json:"username" gorm:"type:varchar(100);unique_index"`
	Email       string     `json:"email" gorm:"type:varchar(100);unique_index"`
	ContactInfo string     `json:"contactInfo"`
	Notes       string     `json:"notes"`
	Password    string     `json:"password,omitempty"`
	Role        roles.Role `json:"role"`
	IsActive    bool       `json:"IsActive"`
}

func (adv Advertiser) Disable() {
	adv.IsActive = false
}

func (adv Advertiser) Activate() {
	adv.IsActive = true
}

func (adv Advertiser) GetHashedPassword() string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adv.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
