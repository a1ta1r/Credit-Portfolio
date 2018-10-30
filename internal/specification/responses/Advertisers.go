package responses

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type AllAdvertisers struct {
	Status string `example:"OK" json:"status"`
	Count int `example:"0" json:"count"`
	Advertisers []entities.Advertiser `example:"[]" json:"advertisers"`
}

type OneAdvertiser struct {
	Advertiser entities.Advertiser `json:"advertiser"`
}
