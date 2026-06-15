package main

import "fmt"

const NMAX int = 100

type Peserta struct {
	id, nama      string
	nilai, durasi int
}

type TabPeserta [NMAX]Peserta

func main() {
	var N int
	var arrPeserta TabPeserta
	fmt.Scan(&N)
	inputPeserta(&arrPeserta, N)
	insertionSort(&arrPeserta, N)
	cetakPeserta(arrPeserta, N)
	fmt.Println("Peserta terbaik:")
	fmt.Println(arrPeserta[0].id, arrPeserta[0].nama, arrPeserta[0].nilai, arrPeserta[0].durasi)
}

func inputPeserta(A *TabPeserta, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i].id, &A[i].nama, &A[i].nilai, &A[i].durasi)
	}
}

func insertionSort(A *TabPeserta, n int) {
	var i, j int
	var temp Peserta

	for i = 1; i < n; i++ {
		j = i
		temp = A[j]
		for j > 0 && (temp.nilai > A[j-1].nilai || (temp.nilai == A[j-1].nilai && temp.durasi < A[j-1].durasi)) {
			A[j] = A[j-1]
			j--
		}
		A[j] = temp
	}
}

func cetakPeserta(A TabPeserta, n int) {
	var i int
	fmt.Print("\nData setelah diurutkan:\n")
	for i = 0; i < n; i++ {
		fmt.Println(A[i].id, A[i].nama, A[i].nilai, A[i].durasi)
	}
}
