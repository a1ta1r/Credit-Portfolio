package responses

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type BannersByAds struct {
	Status  string            `example:"OK" json:"status"`
	Count   int               `example:"0" json:"count"`
	Banners []entities.Banner `example:"[]" json:"banners"`
}

type OneBanner struct {
	Banner entities.Banner `json:"banner"`
}
