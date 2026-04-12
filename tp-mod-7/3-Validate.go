package main

import "fmt"

func main() {
	var konstan int
	fmt.Scan(&konstan)
	validateNumber(konstan)
}

func validateNumber(x int) {
	var tiga, lima, dua bool
	var melanggar bool

	for x != 0 {
		if x%10 == 3 {
			tiga = true
		}

		if x%10 == 5 {
			lima = true
		}

		if x%10 == 2 {
			dua = true
		}

		if tiga && x%10 == 7 || lima && x%10 == 0 || dua && x%10 == 6 {
			melanggar = true
		}

		x = x / 10
	}

	if melanggar {
		fmt.Println("BILANGAN TIDAK VALID")
	} else {
		fmt.Println("BILANGAN VALID")
	}
}
