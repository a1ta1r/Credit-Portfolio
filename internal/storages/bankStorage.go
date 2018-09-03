package storages

import "github.com/jinzhu/gorm"

type bankStorage struct {
	DB gorm.DB
}
