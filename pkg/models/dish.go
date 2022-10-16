package models

import "time"

type Dish struct {
	Id           uint   `gorm:"primaryKey"`
	Name         string `gorm:"index"`
	Price        float32
	Description  string
	Type         string
	Availability bool
	Outlet       uint      `gorm:"index"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
