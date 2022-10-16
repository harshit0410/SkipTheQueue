package models

import "time"

type Order struct {
	Id          uint `gorm:"primaryKey"`
	OutletId    uint
	Status      string
	Type        string
	FinalAmount float32
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
