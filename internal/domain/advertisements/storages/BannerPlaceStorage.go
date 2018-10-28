package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/jinzhu/gorm"
)

type BannerPlaceStorage struct {
	db gorm.DB
}

func (as BannerPlaceStorage) GetBannerPlace(id uint) (models.BannerPlace, error) {
	var bannerPlace models.BannerPlace
	err := as.db.First(&bannerPlace, id).Error
	return bannerPlace, err
}

func (as BannerPlaceStorage) GetBannerPlaces() ([]models.BannerPlace, error) {
	var bannerPlaces []models.BannerPlace
	err := as.db.Find(&bannerPlaces).Error
	return bannerPlaces, err
}

func (as BannerPlaceStorage) CreateBannerPlace(bannerPlace models.BannerPlace) error {
	return as.db.Create(bannerPlace).Error
}

func (as BannerPlaceStorage) UpdateBannerPlace(bannerPlace models.BannerPlace) error {
	return as.db.Save(bannerPlace).Error
}

func (as BannerPlaceStorage) DeleteBannerPlace(bannerPlace models.BannerPlace) error {
	return as.db.Delete(bannerPlace, bannerPlace.ID).Error
}
