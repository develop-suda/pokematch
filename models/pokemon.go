package models

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Name string `gorm:"not null"`
	// 最大精度と最大位取りを設定
	Height float32 `gorm:"type:numeric(3,1);not null"`
	Weight float32 `gorm:"type:numeric(4,1);not null"`
}
