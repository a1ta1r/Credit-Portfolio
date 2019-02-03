package responses

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type AllBannersPlaces struct {
	Status       string                 `example:"OK" json:"status"`
	Count        int                    `example:"0" json:"count"`
	BannerPlaces []entities.BannerPlace `example:"[]" json:"bannerPlaces"`
}

type OneBannerPlace struct {
	BannerPlace entities.BannerPlace `json:"bannerPlace"`
}
