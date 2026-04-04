package main

import "fmt"

func main() {
	var start, nHari int
	fmt.Scan(&start, &nHari)
	fmt.Println(hariKerja(start, nHari))
}

func hariKerja(hariMulai, jumlahHari int) int {
	var temp int
	if jumlahHari == 1 {
		temp = 1
		if hariMulai == 6 || hariMulai == 7 {
			temp = 0
		}
		return temp
	} else {
		temp = 1

		if hariMulai == 6 {
			temp = 0
		} else if hariMulai == 7 {
			temp = 0
			hariMulai = 0
		}

		return hariKerja(hariMulai+1, jumlahHari-1) + temp
	}
}
