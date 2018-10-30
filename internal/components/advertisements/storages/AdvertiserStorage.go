package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	"github.com/jinzhu/gorm"
)

type AdvertiserStorage struct {
	DB gorm.DB
}

func (as AdvertiserStorage) GetAdvertiser(id uint) (entities.Advertiser, error) {
	var advertiser entities.Advertiser
	err := as.DB.First(&advertiser, id).Error
	return advertiser, err
}

func (as AdvertiserStorage) GetAdvertisers() ([]entities.Advertiser, error) {
	var advertisers []entities.Advertiser
	err := as.DB.Find(&advertisers).Error
	return advertisers, err
}

func (as AdvertiserStorage) CreateAdvertiser(advertiser *entities.Advertiser) error {
	return as.DB.Create(&advertiser).Error
}

func (as AdvertiserStorage) UpdateAdvertiser(advertiser entities.Advertiser) error {
	return as.DB.Save(&advertiser).Error
}

func (as AdvertiserStorage) DeleteAdvertiser(advertiser entities.Advertiser) error {
	return as.DB.Delete(&advertiser, advertiser.ID).Error
}
