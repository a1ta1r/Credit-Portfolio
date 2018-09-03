package storages

import "github.com/jinzhu/gorm"

type expenseStorage struct {
	DB gorm.DB
}
