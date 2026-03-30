package main

import "fmt"

func main() {
	var total int
	total = transaksi(1, 0)
	fmt.Printf("Total Poin: %d\n", total)
}

func hitungPoin(n int) int {
	if n < 50000 {
		return 0
	} else if n < 100000 {
		return 5
	} else if n < 200000 {
		return 10
	} else {
		return 20
	}
}

func transaksi(no int, totalPoin int) int {
	var belanja int
	var poin int

	fmt.Scan(&belanja)

	if belanja == 0 {
		return totalPoin
	} else {
		poin = hitungPoin(belanja)
		fmt.Printf("Transaksi %d: Belanja Rp %d, Poin: %d\n", no, belanja, poin)
		return transaksi(no+1, totalPoin+poin)
	}
}
