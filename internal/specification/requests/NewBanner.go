package requests

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type NewBanner struct {
	AdvertisementID     uint   `json:"advertisementId" binding:"required"`
	PictureUrl          string `json:"pictureUrl" binding:"required"`
	Text                string `json:"text" binding:"required"`
	UniqueViewsRequired uint   `json:"uniqueViewsRequired" binding:"required"`
	AdvertisementLink   string `json:"advertisementLink" binding:"required"`
	BannerPlaceID       uint   `json:"bannerPlaceId"  binding:"required"`
	IsVisible           bool   `json:"isVisible" binding:"required"`
}

func (nb NewBanner) ToBanner() entities.Banner {
	return entities.Banner{
		AdvertisementID:     nb.AdvertisementID,
		IsVisible:           nb.IsVisible,
		BannerPlaceID:       nb.BannerPlaceID,
		PictureUrl:          nb.PictureUrl,
		Text:                nb.Text,
		UniqueViewsRequired: nb.UniqueViewsRequired,
		AdvertisementLink:   nb.AdvertisementLink,
	}
}
