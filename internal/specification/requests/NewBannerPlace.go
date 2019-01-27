package requests

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
)

type NewBannerPlace struct {
	PricePerView float64 `json:"pricePerView" binding:"required"`
	Description  string  `json:"description" binding:"required"`
}

func (nb NewBannerPlace) ToBanner() entities.BannerPlace {
	return entities.BannerPlace{
		PricePerView: nb.PricePerView,
		Description:  nb.Description,
	}
}
