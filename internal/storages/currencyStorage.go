package storages

import "github.com/jinzhu/gorm"

type currencyStorage struct {
	DB gorm.DB
}
