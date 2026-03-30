package main

import "fmt"

func main() {
	var nomorHP string
	fmt.Scan(&nomorHP)
	if validasiNomor(nomorHP) {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}
}

func hitungDigitRek(nomor string, posisi int) int {
	if nomor[posisi:] == "" {
		return 1
	} else {
		return hitungDigitRek(nomor, posisi+1) + 1
	}
}

func awalanValid(nomor string) bool {
	if (nomor[0] == '0' && nomor[1] == '8') {
		return true
	} else if (nomor[0] == '+' && nomor[1] == '6' && nomor[2] == '2') {
		return true
	} else {
		return false
	}
}

func validasiNomor(nomor string) bool {
	var jumlahDigitSesuai bool
	var awalan bool

	awalan = awalanValid(nomor)

	if nomor[0] == '+' {
		nomor = nomor[1:]
	}

	jumlahDigitSesuai = hitungDigitRek(nomor, 3) >= 9 && hitungDigitRek(nomor, 3) <= 12

	if jumlahDigitSesuai && awalan {
		return true
	} else {
		return false
	}
}
