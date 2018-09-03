package storages

import "github.com/jinzhu/gorm"

type paymentStorage struct {
	DB gorm.DB
}
