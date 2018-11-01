package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	"github.com/jinzhu/gorm"
)

type AdvertisementStorage struct {
	DB gorm.DB
}

func (as AdvertisementStorage) GetAdvertisement(id uint) (entities.Advertisement, error) {
	var advertisement entities.Advertisement
	err := as.DB.First(&advertisement, id).Error
	return advertisement, err
}

func (as AdvertisementStorage) GetAdvertisements() ([]entities.Advertisement, error) {
	var advertisements []entities.Advertisement
	err := as.DB.Find(&advertisements).Error
	return advertisements, err
}

func (as AdvertisementStorage) GetAdvertisementsByAdvertiser(id uint) ([]entities.Advertisement, error) {
	var advertisements []entities.Advertisement
	err := as.DB.Where("advertiser_id = ?", id).Find(&advertisements).Error
	return advertisements, err
}

func (as AdvertisementStorage) CreateAdvertisement(advertisement *entities.Advertisement) error {
	return as.DB.Create(advertisement).Error
}

func (as AdvertisementStorage) UpdateAdvertisement(advertisement *entities.Advertisement) error {
	return as.DB.Save(advertisement).Error
}

func (as AdvertisementStorage) DeleteAdvertisement(advertisement entities.Advertisement) error {
	return as.DB.Delete(advertisement, advertisement.ID).Error
}
