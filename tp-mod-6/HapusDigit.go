package main

import "fmt"

func main() {
	var n, nDigit, target int

	fmt.Scan(&n)

	nDigit = jumlahDigit(n)

	if nDigit <= 2 {
		fmt.Println("INPUT ERROR")
	} else if nDigit%2 == 0 {
		target = nDigit / 2
		fmt.Println(hapusGenap(n, target))
	} else {
		target = nDigit/2 + 1
		fmt.Println(hapusGanjil(n, target))
	}
}

func jumlahDigit(n int) int {
	if n < 10 {
		return 1
	} else {
		return jumlahDigit(n/10) + 1
	}
}

func hapusGanjil(n, del int) int {
	if del == 1 {
		return n / 10
	} else {
		return hapusGanjil(n/10, del-1)*10 + n%10
	}
}

func hapusGenap(n, del int) int {
	if del == 1 {
		return n / 100
	} else {
		return hapusGenap(n/10, del-1)*10 + n%10
	}
}
