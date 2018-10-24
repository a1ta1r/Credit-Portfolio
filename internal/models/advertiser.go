package models

import "time"

type Advertiser struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
	Username    string    `json:"username" gorm:"type:varchar(100);unique_index"`
	Email       string    `json:"email" gorm:"type:varchar(100);unique_index"`
	ContactInfo string    `json:"contactInfo"`
	Notes       string    `json:"notes"`
	Password    string    `json:"password,omitempty"`
	Role        Role      `json:"role"`
	isActive    bool
}

func CreateAdvertiser(username string, email string, password string) Advertiser {
	return Advertiser{
		Username: username,
		Email:    email,
		Password: password,
		Role:     Ads,
	}
}

func (adv Advertiser) Disable() {
	adv.isActive = false
}

func (adv Advertiser) Activate() {
	adv.isActive = true
}
