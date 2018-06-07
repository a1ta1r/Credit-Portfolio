package models

import "github.com/jinzhu/gorm"

type PaymentType struct {
	gorm.Model
	name string
}
