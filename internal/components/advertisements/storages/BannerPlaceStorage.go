package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	"github.com/jinzhu/gorm"
)

type BannerPlaceStorage struct {
	DB gorm.DB
}

func (as BannerPlaceStorage) GetBannerPlace(id uint) (entities.BannerPlace, error) {
	var bannerPlace entities.BannerPlace
	err := as.DB.First(&bannerPlace, id).Error
	return bannerPlace, err
}

func (as BannerPlaceStorage) GetBannerPlaces() ([]entities.BannerPlace, error) {
	var bannerPlaces []entities.BannerPlace
	err := as.DB.Find(&bannerPlaces).Error
	return bannerPlaces, err
}

func (as BannerPlaceStorage) CreateBannerPlace(bannerPlace *entities.BannerPlace) error {
	return as.DB.Create(bannerPlace).Error
}

func (as BannerPlaceStorage) UpdateBannerPlace(bannerPlace *entities.BannerPlace) error {
	return as.DB.Save(bannerPlace).Error
}

func (as BannerPlaceStorage) DeleteBannerPlace(bannerPlace entities.BannerPlace) error {
	return as.DB.Delete(bannerPlace, bannerPlace.ID).Error
}
