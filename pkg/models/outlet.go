package pkg

import "time"

type Outlet struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"index"`
	Address   string
	PinCode   string
	City      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
