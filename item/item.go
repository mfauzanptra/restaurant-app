package item

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Item struct {
	Id    uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name  string
	Stock int
	Price float64
}

type ItemAuth struct {
	DB *gorm.DB
}

func (ia *ItemAuth) CheckDuplicateItem(name string) error {
	existed := 0
	ia.DB.Raw("SELECT COUNT(*) FROM items i WHERE i.name = ?", name).Scan(&existed)
	if existed >= 1 {
		return errors.New("duplicated item")
	}
	return nil
}
func (ia *ItemAuth) AddItem(newItem Item) error {
	err := ia.CheckDuplicateItem(newItem.Name)
	if err != nil {
		return err
	}
	err = ia.DB.Create(&newItem).Error
	if err != nil {
		log.Println("error creating new item: ", err)
		return err
	}
	return nil
}
func (ia *ItemAuth) GetItems() ([]Item, error) {
	items := []Item{}
	err := ia.DB.Raw("SELECT * FROM items").Scan(&items).Error
	if err != nil {
		log.Println("error getting items: ", err)
		return nil, err
	}
	return items, nil
}
