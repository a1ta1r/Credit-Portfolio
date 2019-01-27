package requests

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
)

type UpdateBannerPlace struct {
	PricePerView *float64 `json:"pricePerView"`
	Description  *string  `json:"description"`
}

func (nb UpdateBannerPlace) ToBannerPlace(bannerPlace entities.BannerPlace) entities.BannerPlace {
	if nb.PricePerView != nil {
		bannerPlace.PricePerView = *nb.PricePerView
	}
	if nb.Description != nil {
		bannerPlace.Description = *nb.Description
	}

	return bannerPlace
}
