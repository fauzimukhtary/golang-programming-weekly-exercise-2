package main

import "fmt"

type kota struct {
	nama     string
	penduduk int
	laporan  float64
}

type tabKota [20]kota

func main() {
	var kota tabKota
	var i, jlhKota, populasiAwal int
	fmt.Scan(&jlhKota, &populasiAwal)

	for i = 0; i < jlhKota; i++ {
		fmt.Scan(&kota[i].nama)
		kota[i].penduduk = populasiAwal / 10
		populasiAwal -= kota[i].penduduk
		kota[i].laporan = float64(kota[i].penduduk%1000) * 0.025
	}

	for i = 0; i < jlhKota; i++ {
		if kota[i].laporan > 5.0 {
			fmt.Printf("%v dengan tingkat kejahatan: %.2f%%\n", kota[i].nama, kota[i].laporan)
		}
	}
}
