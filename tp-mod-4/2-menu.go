package main

import "fmt"

const NMAX int = 999

type menu struct {
	name         string
	price, stock int
	available    string
}

type arrMenu [NMAX]menu

func main() {
	var n int
	var array arrMenu
	fmt.Scan(&n)
	inputDataMenu(&array, n)
	printDataMenu(array, n)
}

func inputDataMenu(arr *arrMenu, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("\nNama Menu: ")
		fmt.Scan(&arr[i].name)
		fmt.Print("Harga Menu: ")
		fmt.Scan(&arr[i].price)
		fmt.Print("Stock: ")
		fmt.Scan(&arr[i].stock)
		if arr[i].stock == 0 {
			arr[i].available = "Kosong"
		} else {
			arr[i].available = "Tersedia"
		}
	}
}

func printDataMenu(arr arrMenu, n int) {
	var i int
	fmt.Println()
	fmt.Print("Nama Menu\tHarga (Rp)\tStok (pcs)\tStatus Kesediaan\n")
	for i = 0; i < n; i++ {
		fmt.Printf("%v\t", arr[i].name)
		if len(arr[i].name) < 8 {
			fmt.Print("\t")
		}

		fmt.Printf("Rp%d\t\t", arr[i].price)

		fmt.Printf("%d\t\t", arr[i].stock)
		fmt.Printf("%v\n", arr[i].available)
	}
	fmt.Println()
}
