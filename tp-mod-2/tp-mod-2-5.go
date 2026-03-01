package main

import "fmt"

func tentukanKategori(umur int, beratSaatIni float64) {
	var ideal float64
	var kategori string

	ideal = 2.20462 * (float64(umur)*2 + 8)

	if beratSaatIni < 0.9*ideal {
		kategori = "kurang"
	} else if beratSaatIni <= 1.1*ideal {
		kategori = "ideal"
	} else {
		kategori = "obesitas"
	}

	fmt.Printf("%.2f lb\n", ideal)
	fmt.Printf("Kategori: %v\n", kategori)
}

func main() {
	var umur int
	var berat float64

	fmt.Scan(&umur, &berat)
	tentukanKategori(umur, berat)
}
