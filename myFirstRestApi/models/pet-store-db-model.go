package models

import "gorm.io/gorm"

type PetStoreDb struct {
	gorm.Model
	Name string
}