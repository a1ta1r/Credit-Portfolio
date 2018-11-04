package requests

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type UpdateAdvertiser struct {
	Username    *string `json:"username" binding:"min=3"`
	Email       *string `json:"email" binding:"email"`
	ContactInfo *string `json:"contactInfo"`
	Notes       *string `json:"notes"`
	IsActive    *bool    `json:"isActive"`
}

func (ua UpdateAdvertiser) ToAdvertiser(advertiser entities.Advertiser) entities.Advertiser {
	if ua.Username != nil {
		advertiser.Username = *ua.Username
	}
	if ua.Email != nil {
		advertiser.Email = *ua.Email
	}
	if ua.ContactInfo != nil {
		advertiser.ContactInfo = *ua.ContactInfo
	}
	if ua.Notes != nil {
		advertiser.Notes = *ua.Notes
	}
	if ua.IsActive != nil {
		advertiser.IsActive = *ua.IsActive
	}
	return advertiser
}
