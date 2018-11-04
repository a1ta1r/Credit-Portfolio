package requests

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type NewAdvertisement struct {
	AdvertiserID uint   `json:"advertiserId" binding:"required"`
	IsActive     bool   `json:"isActive" binding:"required"`
	Title        string `json:"title" binding:"required"`
}

func (na NewAdvertisement) ToAdvertisement() entities.Advertisement {
	return entities.Advertisement{
		AdvertiserID: na.AdvertiserID,
		IsActive:     na.IsActive,
		Title:        na.Title,
	}
}
