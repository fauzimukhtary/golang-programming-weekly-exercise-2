package main

import "fmt"

func main() {
	var angka, nDigit int

	fmt.Scan(&angka)
	nDigit = countDigit(angka)

	if nDigit < 9 || nDigit > 9 {
		fmt.Println(false)
	} else {
		fmt.Println(validSolution(angka))
	}
}

func validSolution(n int) bool {
	if n == 0 {
		return true
	} else if (hitungKemunculan(n, n%10) > 1) || (n%10 == 0 && n > 9) {
		return false
	} else {
		return validSolution(n / 10)
	}
}

func hitungKemunculan(n, d int) int {
	if n == 0 {
		return 0
	} else {
		if n%10 == d {
			return hitungKemunculan(n/10, d) + 1
		} else {
			return hitungKemunculan(n/10, d)
		}
	}
}

func countDigit(n int) int {
	if n == 0 {
		return 0
	} else {
		return countDigit(n/10) + 1
	}
}
