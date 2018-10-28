package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/jinzhu/gorm"
)

type AdvertiserStorage struct {
	db gorm.DB
}

func (as AdvertiserStorage) GetAdvertiser(id uint) (models.Advertiser, error) {
	var advertiser models.Advertiser
	err := as.db.First(&advertiser, id).Error
	return advertiser, err
}

func (as AdvertiserStorage) GetAdvertisers() ([]models.Advertiser, error) {
	var advertisers []models.Advertiser
	err := as.db.Find(&advertisers).Error
	return advertisers, err
}

func (as AdvertiserStorage) CreateAdvertiser(advertiser models.Advertiser) error {
	return as.db.Create(&advertiser).Error
}

func (as AdvertiserStorage) UpdateAdvertiser(advertiser models.Advertiser) error {
	return as.db.Save(&advertiser).Error
}

func (as AdvertiserStorage) DeleteAdvertiser(advertiser models.Advertiser) error {
	return as.db.Delete(&advertiser, advertiser.ID).Error
}
