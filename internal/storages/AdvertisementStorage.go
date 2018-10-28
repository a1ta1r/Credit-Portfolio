package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/jinzhu/gorm"
)

type AdvertisementStorage struct {
	db gorm.DB
}

func (as AdvertisementStorage) GetAdvertisement(id uint) (models.Advertisement, error) {
	var advertisement models.Advertisement
	err := as.db.First(&advertisement, id).Error
	return advertisement, err
}

func (as AdvertisementStorage) GetAdvertisements() ([]models.Advertisement, error) {
	var advertisements []models.Advertisement
	err := as.db.Find(&advertisements).Error
	return advertisements, err
}

func (as AdvertisementStorage) GetAdvertisementsByAdvertiser(id uint) ([]models.Advertisement, error) {
	var advertisements []models.Advertisement
	err := as.db.Where("advertiser_id = ?", id).Find(&advertisements).Error
	return advertisements, err
}

func (as AdvertisementStorage) CreateAdvertisement(advertisement models.Advertisement) error {
	return as.db.Create(advertisement).Error
}

func (as AdvertisementStorage) UpdateAdvertisement(advertisement models.Advertisement) error {
	return as.db.Save(advertisement).Error
}

func (as AdvertisementStorage) DeleteAdvertisement(advertisement models.Advertisement) error {
	return as.db.Delete(advertisement, advertisement.ID).Error
}
