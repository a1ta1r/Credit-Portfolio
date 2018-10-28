package common

import (
	as "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/storages"
	ls "github.com/a1ta1r/Credit-Portfolio/internal/components/loans/storages"
	"github.com/jinzhu/gorm"
)

type StorageContainer struct {
	UserStorage          ls.UserStorage
	AdvertiserStorage    as.AdvertiserStorage
	AdvertisementStorage as.AdvertisementStorage
	BannerStorage        as.BannerStorage
	BannerPlaceStorage   as.BannerPlaceStorage
}

func NewStorageContainer(db gorm.DB) StorageContainer {
	return StorageContainer{
		UserStorage:          ls.UserStorage{DB: db},
		AdvertiserStorage:    as.AdvertiserStorage{DB: db},
		AdvertisementStorage: as.AdvertisementStorage{DB: db},
		BannerStorage:        as.BannerStorage{DB: db},
		BannerPlaceStorage:   as.BannerPlaceStorage{DB: db},
	}
}
