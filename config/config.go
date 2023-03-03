package config

import (
	"fmt"
	"os"
	"restaurant-app/item"
	"restaurant-app/order"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

func ReadConfig() *Config {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("error read env file: ", err.Error())
		return nil
	}

	var res Config
	res.DBUser = os.Getenv("DBUSER")
	res.DBPass = os.Getenv("DBPASS")
	res.DBHost = os.Getenv("DBHOST")
	readData := os.Getenv("DBPORT")
	res.DBPort, err = strconv.Atoi(readData)
	if err != nil {
		fmt.Println("error convert: ", err.Error())
		return nil
	}
	res.DBName = os.Getenv("DBNAME")
	return &res
}

func ConnectSQL(c Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection error: ", err.Error())
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(item.Item{})
	db.AutoMigrate(order.Order{})
	db.AutoMigrate(order.OrderItem{})
}
