package requests

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type NewAdvertiser struct {
	Username    string `json:"username" binding:"required,min=3"`
	Email       string `json:"email" binding:"required,email"`
	ContactInfo string `json:"contactInfo"`
	Notes       string `json:"notes"`
	Password    string `json:"password,omitempty" binding:"required,min=3"`
	IsActive    bool   `json:"isActive" binding:"required"`
}

func (nr NewAdvertiser) ToAdvertiser() entities.Advertiser {
	return entities.Advertiser{
		Username:nr.Username,
		Email:nr.Email,
		ContactInfo:nr.ContactInfo,
		Notes:nr.Notes,
		Password:nr.Password,
		IsActive:nr.IsActive,
	}
}
