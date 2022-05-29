package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name string
	Address string
	Country string
}