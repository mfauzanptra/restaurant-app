package order

import (
	"log"
	"restaurant-app/item"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TotalPrice float64
}

type OrderItem struct {
	OrderId  uint
	Order    Order
	ItemId   uint
	Item     item.Item
	Quantity int
}

type Cart struct {
	ItemId uint
	Qty    int
}

type OrderAuth struct {
	DB *gorm.DB
}

func (oa *OrderAuth) CalculateTotalPrices(cart []Cart) float64 {
	totalPrice := 0.0
	for _, v := range cart {
		price := 0.0
		oa.DB.Raw("SELECT price FROM items WHERE id = ?", v.ItemId).Scan(&price)
		totalPrice += price * float64(v.Qty)
	}
	return totalPrice
}

func (oa *OrderAuth) CreateOrder(cart []Cart) error {
	order := Order{}
	ordertItems := []OrderItem{}

	tx := oa.DB.Begin()

	order.TotalPrice = oa.CalculateTotalPrices(cart)

	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		log.Printf("Error creating order")
		return err
	}

	for _, v := range cart {
		orderItem := OrderItem{
			OrderId:  order.ID,
			ItemId:   v.ItemId,
			Quantity: v.Qty,
		}
		ordertItems = append(ordertItems, orderItem)
	}

	err = tx.Create(&ordertItems).Error
	if err != nil {
		tx.Rollback()
		log.Println("Error creating orderitems ", err)
		return err
	}
	tx.Commit()
	return nil
}
