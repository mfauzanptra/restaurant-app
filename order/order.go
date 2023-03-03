package order

import (
	"restaurant-app/item"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ItemId     uint
	Item       item.Item
	TotalPrice float64
}

type OrderItem struct {
	OrderId  uint
	Order    Order
	ItemId   uint
	Item     item.Item
	Quantity int
}

type OrderAuth struct {
	DB *gorm.DB
}
