package entities

import "time"

type Advertisement struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	Advertiser   Advertiser `json:"-"`
	AdvertiserID uint       `json:"advertiserId"`
	IsActive     bool       `json:"IsActive"`
	Banners      []Banner   `json:"banners"`
}
