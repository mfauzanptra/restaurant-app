package main

import (
	"restaurant-app/config"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	var cfg = config.ReadConfig()
	var db = config.ConnectSQL(*cfg)
	config.Migrate(db)

}
