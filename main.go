package main

import (
	"bufio"
	"fmt"
	"os"
	"restaurant-app/cart"
	"restaurant-app/config"
	"restaurant-app/item"
	"restaurant-app/order"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var cfg = config.ReadConfig()
	var db = config.ConnectSQL(*cfg)
	config.Migrate(db)
	var itemAuth = item.ItemAuth{DB: db}
	var orderAuth = order.OrderAuth{DB: db}

	uCart := []cart.Cart{}
	option := 0
	for option != 9 {
		fmt.Println("\n=======================")
		fmt.Println("Rumah Makan ABC")
		fmt.Println("1. Tambah Menu Makanan")
		fmt.Println("2. Lihat Menu Makanan")
		fmt.Println("3. Buat Pesanan")
		fmt.Println("4. Laporan Penjualan")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan menu: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var item item.Item
			fmt.Println("/nTambah Menu Makanan")
			fmt.Print("Nama Menu: ")
			scanner.Scan()
			item.Name = scanner.Text()
			fmt.Print("Harga: ")
			fmt.Scanln(&item.Price)
			fmt.Print("Stok: ")
			fmt.Scanln(&item.Stock)

			err := itemAuth.AddItem(item)
			if err != nil {
				fmt.Println("\ngagal menambahkan menu: ", err)
			} else {
				fmt.Println("\nsukses menambahkan menu")
			}
		case 2:
			fmt.Println("\nList Menu")
			items, err := itemAuth.GetItems()
			if err != nil {
				fmt.Println("\ngagal menampilkan menu: ", err)
			} else {
				for _, item := range items {
					fmt.Println("- ID: ", item.Id)
					fmt.Println("  Nama Menu: ", item.Name)
					fmt.Println("  Stok: ", item.Stock)
					fmt.Println("  Harga: ", item.Price)
				}
			}
		case 3:
			orderMenu := 0
			for orderMenu != 9 {
				fmt.Println("\nBuat Pesanan")
				fmt.Println("1. Tambah Item")
				fmt.Println("2. Lihat Daftar Pesanan")
				fmt.Println("3. Bayar Pesanan")
				fmt.Println("9. Kembali")
				fmt.Print("Enter an option : ")
				fmt.Scanln(&orderMenu)

				switch orderMenu {
				case 1:
					fmt.Println("\nList Menu")
					items, err := itemAuth.GetItems()
					if err != nil {
						fmt.Println("\ngagal menampilkan menu: ", err)
					} else {
						for _, item := range items {
							fmt.Println("- ID: ", item.Id)
							fmt.Println("  Nama Menu: ", item.Name)
							fmt.Println("  Stok: ", item.Stock)
							fmt.Println("  Harga: ", item.Price)
						}
						tmp := cart.Cart{}
						fmt.Print("Masukkan ID item: ")
						fmt.Scanln(&tmp.ItemId)
						fmt.Print("Jumlah: ")
						fmt.Scanln(&tmp.Qty)
						uCart = append(uCart, tmp)
					}
				case 2:
					fmt.Println("\nDaftar Pesanan")
					for _, item := range uCart {
						tmp := itemAuth.GetItem(item.ItemId)
						fmt.Println("- ID: ", tmp.Id)
						fmt.Println("  Nama Item: ", tmp.Name)
						fmt.Println("  Jumlah: ", item.Qty)
						fmt.Println("  Harga: ", tmp.Price)
					}
					fmt.Printf("Total Harga: %.2f\n", orderAuth.CalculateTotalPrices(uCart))
				case 3:
					if len(uCart) == 0 {
						fmt.Println("Daftar pesananan masih kosong")
					} else {
						fmt.Println("\nDaftar Pesanan")
						for _, item := range uCart {
							tmp := itemAuth.GetItem(item.ItemId)
							fmt.Println("- ID: ", tmp.Id)
							fmt.Println("  Nama Item: ", tmp.Name)
							fmt.Println("  Jumlah: ", item.Qty)
							fmt.Println("  Harga: ", tmp.Price)
						}
						fmt.Printf("Total Harga: %.2f\n", orderAuth.CalculateTotalPrices(uCart))

						fmt.Print("Pesanan sudah sesuai? (y/n): ")
						answer := ""
						fmt.Scanln(&answer)
						if answer == "n" {
							continue
						} else {
							err := orderAuth.CreateOrder(uCart)
							if err != nil {
								fmt.Println("gagal membuat pesanan")
							} else {
								fmt.Println("sukses membuat pesanan")
							}
						}
					}
				}
			}
		}
	}
}
