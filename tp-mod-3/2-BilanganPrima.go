package main

import "fmt"

func isPrima(angka int) bool {
	var i int
	var prime bool

	prime = true
	if angka < 2 {
		prime = false
	} else {
		i = 2
		for i*i <= angka {
			if angka%i == 0 {
				prime = false
			}
			i++
		}
	}

	return prime
}

func sumAllDigit(angka int) int {
	var hasil int

	hasil = 0
	for angka != 0 {
		hasil = hasil + (angka % 10)
		angka = angka / 10
	}

	return hasil
}

func sumUntilBeforeLimit(b1, b2, limit int) int {
	var hasil int
	hasil = b1
	for hasil < limit {
		hasil = hasil + b2
	}

	hasil = hasil - b2
	return hasil
}

func eksperimen(angka int) {
	var penambah, temp, hasil int
	var prima bool

	temp = angka
	if angka < 100 {
		if angka%10 == 0 {
			penambah = angka / 10
		} else {
			penambah = angka % 10
		}
		hasil = sumUntilBeforeLimit(temp, penambah, 1000)
	} else if angka < 1000 {
		penambah = sumAllDigit(angka)
		hasil = sumUntilBeforeLimit(temp, penambah, 10000)
	} else {
		hasil = angka
	}

	fmt.Printf("%d --> %d\n", angka, hasil)

	prima = isPrima(temp)
	if prima {
		fmt.Println("Prima")
	} else {
		fmt.Println("Bukan Prima")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	eksperimen(n)
}
