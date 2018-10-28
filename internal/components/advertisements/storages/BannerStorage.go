package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	"github.com/jinzhu/gorm"
)

type BannerStorage struct {
	DB gorm.DB
}

func (as BannerStorage) GetBanner(id uint) (entities.Banner, error) {
	var banner entities.Banner
	err := as.DB.First(&banner, id).Error
	return banner, err
}

func (as BannerStorage) GetBannersByAdvertisement(id uint) ([]entities.Banner, error) {
	var banners []entities.Banner
	err := as.DB.Where(entities.Banner{AdvertisementID: id}).Find(&banners).Error
	return banners, err
}

func (as BannerStorage) GetBanners() ([]entities.Banner, error) {
	var banners []entities.Banner
	err := as.DB.Find(&banners).Error
	return banners, err
}

func (as BannerStorage) CreateBanner(banner entities.Banner) error {
	return as.DB.Create(banner).Error
}

func (as BannerStorage) UpdateBanner(banner entities.Banner) error {
	return as.DB.Save(banner).Error
}

func (as BannerStorage) DeleteBanner(banner entities.Banner) error {
	return as.DB.Delete(banner, banner.ID).Error
}
