package entities

import "time"

type Banner struct {
	ID                  uint          `gorm:"primary_key" json:"id"`
	CreatedAt           time.Time     `json:"createdAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
	PictureUrl          string        `json:"pictureUrl"`
	Text                string        `json:"text"`
	UniqueViewsRequired uint          `json:"uniqueViewsRequired"`
	Views               uint          `json:"views"`
	AdvertisementLink   string        `json:"advertisementLink"`
	UniqueViews         uint          `json:"uniqueViews"`
	BannerPlace         BannerPlace   `json:"-" gorm:"foreignkey:BannerPlaceID"`
	BannerPlaceID       uint          `json:"bannerPlaceId"`
	IsVisible           bool          `json:"isVisible"`
	Advertisement       Advertisement `json:"-" gorm:"foreignkey:AdvertisementID"`
	AdvertisementID     uint          `json:"advertisementId"`
}

func (b Banner) GetBannerPrice() float64 {
	return float64(b.UniqueViewsRequired) * b.BannerPlace.PricePerView
}
