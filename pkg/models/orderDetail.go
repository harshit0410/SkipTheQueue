package models

type OrderDetails struct {
	Id         uint `gorm:"primaryKey"`
	OrderId    uint
	DishId     uint
	Quantity   uint
	Price      float32
	TotalPrice float32
}
