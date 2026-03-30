package main

import "fmt"

func main() {
	var prim, digitSama, kelipatanTiga bool
	var angka, nDigit, blgnTengah int

	fmt.Scan(&angka)
	nDigit = countDigit(angka)
	blgnTengah = ambilBilanganTengah(angka)

	prim = isPrima(blgnTengah, 2)
	digitSama = checkFirstLast(angka, nDigit)
	kelipatanTiga = sumDigit(angka)%3 == 0

	switch {
	case kelipatanTiga && digitSama && prim:
		fmt.Println("Spesial")
	case !digitSama:
		fmt.Println("Beresiko")
	case !prim:
		fmt.Println("Biasa")
	case !kelipatanTiga:
		fmt.Println("Tidak Layak Terbang")
	}
}

func isPrima(n, i int) bool {
	if n < 2 {
		return false
	}

	if i >= n {
		return true
	}

	if n%i == 0 {
		return false
	}

	return isPrima(n, i+1)
}

func ambilBilanganTengah(n int) int {
	if n < 100 {
		return 0
	} else {
		return ambilBilanganTengah(n/10)*10 + (n/10)%10
	}
}

func countDigit(n int) int {
	if n < 10 {
		return 1
	} else {
		return countDigit(n/10) + 1
	}
}

func getNumber(n, i int) int {
	if i == 1 {
		return n % 10
	} else {
		return getNumber(n/10, i-1)
	}
}

func checkFirstLast(n, nDigit int) bool {
	return getNumber(n, 1) == getNumber(n, nDigit)
}

func sumDigit(n int) int {
	if n < 10 {
		return n
	} else {
		return sumDigit(n/10) + n%10
	}
}
