package item

import (
	"gorm.io/gorm"
)

type Item struct {
	Id    uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name  string
	Price float64
}

type ItemAuth struct {
	DB *gorm.DB
}

// func (ia *ItemAuth) CheckDuplicateItem(name string) error {
// 	existed := 0
// 	ia.db.Raw("SELECT COUNT(*) FROM items i WHERE p.product_name = ? AND user_id = ?", newProduct.ProductName, userId).Scan(&existed)
// 	if existed >= 1 {
// 		return errors.New("duplicated product on name")
// 	}
// }
// func (ia *ItemAuth) AddItem(name string, price float64) error {

// }
