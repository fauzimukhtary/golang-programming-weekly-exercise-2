package main

import "fmt"

type item struct {
	nama  string
	harga int
}

type arrItem [20]item

func main() {
	var benda arrItem
	var i, n int
	var min, max item
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Print("Nama Item: ")
		fmt.Scan(&benda[i].nama)
		fmt.Print("Harga Item: ")
		fmt.Scan(&benda[i].harga)
		fmt.Println()
	}

	min = benda[0]
	max = benda[0]

	fmt.Println("Nama Item\t\tHarga")
	for i = 0; i < n; i++ {
		fmt.Println(benda[i].nama, "\t\t", benda[i].harga)

		if benda[i].harga > max.harga {
			max = benda[i]
		}

		if benda[i].harga < min.harga {
			min = benda[i]
		}
	}
	fmt.Println("Item termahal", max.nama)
	fmt.Println("Item termurah", min.nama)
}
