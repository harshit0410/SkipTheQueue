package pkg

import "time"

type Dish struct {
	Id           uint   `gorm:"primaryKey"`
	Name         string `gorm:"index"`
	Description  string
	Type         string
	Availability bool
	Outlet       uint
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
