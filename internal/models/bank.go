package models

import "github.com/jinzhu/gorm"

type Bank struct {
	gorm.Model
	name string
}
