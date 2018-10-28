package models

import "time"

type Advertisement struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    time.Time  `json:"deletedAt"`
	Advertiser   Advertiser `json:"-"`
	AdvertiserID uint       `json:"advertiserId"`
	IsActive     bool       `json:"IsActive"`
	Banners      []Banner   `json:"banners"`
}
