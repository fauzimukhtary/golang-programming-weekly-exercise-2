package main

import "fmt"

const NMAX int = 99

type Penduduk struct {
	nama     string
	thnLahir int
	tpLahir  string
}

type tabPenduduk [NMAX]Penduduk

func main() {
	var N, res int
	var daftarPenduduk tabPenduduk
	var target int
	fmt.Scan(&N)
	inputPenduduk(&daftarPenduduk, N)
	fmt.Scan(&target)
	res = cariPenduduk(daftarPenduduk, N, target)
	if res == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		fmt.Println(daftarPenduduk[res].nama)
		fmt.Println(daftarPenduduk[res].thnLahir)
		fmt.Println(daftarPenduduk[res].tpLahir)
		fmt.Println("ditemukan di index ke", res)
	}
}

func inputPenduduk(A *tabPenduduk, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i].nama, &A[i].thnLahir, &A[i].tpLahir)
	}
}

func cariPenduduk(A tabPenduduk, n, x int) int {
	var left, right, mid int
	var found int

	left = 0
	right = n - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if A[mid].thnLahir == x {
			found = mid
		} else if x < A[mid].thnLahir {
			right = mid - 1
		} else if x > A[mid].thnLahir {
			left = mid + 1
		}
	}
	return found
}
