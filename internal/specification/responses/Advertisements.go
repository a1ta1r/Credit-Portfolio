package responses

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type AllAdvertisements struct {
	Status string `example:"OK" json:"status"`
	Count int `example:"0" json:"count"`
	Advertisements []entities.Advertisement `example:"[]" json:"advertisements"`
}

type OneAdvertisement struct {
	Advertisement entities.Advertisement `json:"Advertisement"`
}

