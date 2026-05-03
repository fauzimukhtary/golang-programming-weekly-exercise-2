package main

import "fmt"

const NMAX int = 50

type Raport struct {
	matkul string
	nilai  int
}

type tabRaport [NMAX]Raport

func main() {
	var N, target int
	var dataRaport tabRaport

	fmt.Scan(&N)
	inputDataRaport(&dataRaport, N)
	fmt.Scan(&target)
	sekuensial(dataRaport, N, target)
}

func inputDataRaport(t *tabRaport, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&t[i].matkul, &t[i].nilai)
	}
}

func sekuensial(t tabRaport, n, X int) {
	var i int
	var found bool

	found = false
	fmt.Println("---RESULT---")
	for i = 0; i < n; i++ {
		if t[i].nilai == X {
			fmt.Println(t[i].matkul, t[i].nilai)
			found = true
		}
	}

	if found {
		fmt.Println("Data Ditemukan!")
	} else {
		fmt.Println("Data tidak ditemukan!")
	}
}
