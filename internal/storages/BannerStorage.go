package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/jinzhu/gorm"
)

type BannerStorage struct {
	db gorm.DB
}

func (as BannerStorage) GetBanner(id uint) (models.Banner, error) {
	var banner models.Banner
	err := as.db.First(&banner, id).Error
	return banner, err
}

func (as BannerStorage) GetBannersByAdvertisement(id uint) ([]models.Banner, error) {
	var banners []models.Banner
	err := as.db.Where(models.Banner{AdvertisementID: id}).Find(&banners).Error
	return banners, err
}

func (as BannerStorage) GetBanners() ([]models.Banner, error) {
	var banners []models.Banner
	err := as.db.Find(&banners).Error
	return banners, err
}

func (as BannerStorage) CreateBanner(banner models.Banner) error {
	return as.db.Create(banner).Error
}

func (as BannerStorage) UpdateBanner(banner models.Banner) error {
	return as.db.Save(banner).Error
}

func (as BannerStorage) DeleteBanner(banner models.Banner) error {
	return as.db.Delete(banner, banner.ID).Error
}
