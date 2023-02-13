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

type OrderType int

const (
	Ordered   OrderType = iota + 1 // EnumIndex = 1
	Preparing                      // EnumIndex = 2
	Ready                          // EnumIndex = 3
)

// String - Creating common behavior - give the type a String function
func (d OrderType) String() string {
	return [...]string{"Ordered", "Preparing", "Ready"}[d-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (d OrderType) EnumIndex() int {
	return int(d)
}
