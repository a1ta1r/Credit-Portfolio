package models

import "github.com/jinzhu/gorm"

type Currency struct {
	gorm.Model
	name    string
	isoCode string
	symbol  string
}
