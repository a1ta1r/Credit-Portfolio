package requests

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type UpdateAdvertiser struct {
	Username    string `json:"username" binding:"min=3"`
	Email       string `json:"email" binding:"email"`
	ContactInfo string `json:"contactInfo"`
	Notes       string `json:"notes"`
	IsActive    bool   `json:"isActive"`
}

func (ua UpdateAdvertiser) ToAdvertiser(advertiser entities.Advertiser) entities.Advertiser {
	advertiser.Username = ua.Username
	advertiser.Email = ua.Email
	advertiser.ContactInfo = ua.ContactInfo
	advertiser.Notes = ua.Notes
	advertiser.IsActive = ua.IsActive
	return advertiser
}
