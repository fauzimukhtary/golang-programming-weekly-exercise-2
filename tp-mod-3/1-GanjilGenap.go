package main

import "fmt"

func reverseAngka(angka int) int {
	var hasilReverse int

	hasilReverse = 0
	for angka != 0 {
		hasilReverse = (hasilReverse * 10) + (angka % 10)
		angka = angka / 10
	}

	return hasilReverse
}

func ambilGabungkanTigaDigit(x int) int {
	var temp, hasil int

	temp = 0
	for temp <= 99 {
		temp = (temp * 10) + (x % 10)
		x = x / 10
	}

	hasil = x*1000 + temp
	return hasil
}

func cekGanjilGenap(angka int) {
	var temp, genap, ganjil int

	temp = angka
	genap = 0
	ganjil = 0
	for temp != 0 {
		if (temp%10)%2 == 0 {
			genap++
		} else {
			ganjil++
		}
		temp = temp / 10
	}

	if ganjil > genap {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}

func prosesPola(angka int) {
	angka = reverseAngka(angka)
	angka = ambilGabungkanTigaDigit(angka)
	fmt.Println(angka)
	cekGanjilGenap(angka)
}

func main() {
	var n int
	fmt.Scan(&n)
	prosesPola(n)
}
