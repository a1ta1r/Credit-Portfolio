package requests

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type UpdateBanner struct {
	AdvertisementID     *uint   `json:"advertisementId"`
	PictureUrl          *string `json:"pictureUrl"`
	Text                *string `json:"text"`
	UniqueViewsRequired *uint   `json:"uniqueViewsRequired"`
	AdvertisementLink   *string `json:"advertisementLink"`
	BannerPlaceID       *uint   `json:"bannerPlaceId" `
	IsVisible           *bool   `json:"isVisible"`
}

func (nb UpdateBanner) ToBanner(banner entities.Banner) entities.Banner {
	if nb.BannerPlaceID != nil {
		banner.BannerPlaceID = *nb.BannerPlaceID
	}
	if nb.IsVisible != nil {
		banner.IsVisible = *nb.IsVisible
	}
	if nb.PictureUrl != nil {
		banner.PictureUrl = *nb.PictureUrl
	}
	if nb.Text != nil {
		banner.Text = *nb.Text
	}
	if nb.UniqueViewsRequired != nil {
		banner.UniqueViewsRequired = *nb.UniqueViewsRequired
	}
	if nb.AdvertisementLink != nil {
		banner.AdvertisementLink = *nb.AdvertisementLink
	}
	return banner
}
