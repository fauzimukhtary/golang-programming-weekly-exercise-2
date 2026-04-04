package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	if isPeak(number) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func isPeak(n int) bool {
	var a1, a2, a3 int

	if n < 100 {
		return false
	}

	a1 = n / 100 % 10
	a2 = n / 10 % 10
	a3 = n % 10

	if a2 > a1 && a2 > a3 {
		return true
	} else {
		return isPeak(n / 10)
	}
}
