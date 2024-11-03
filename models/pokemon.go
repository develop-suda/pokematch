package models

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Name   string  `gorm:"not null"`
	Height float32 `gorm:"not null"`
	Weight float32 `gorm:"not null"`
}
