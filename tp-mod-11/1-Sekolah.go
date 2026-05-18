package main

import "fmt"

const NMAX int = 1001

type tabMurid [NMAX]string

func main() {
	var N, idxRes int
	var target string
	var daftarMurid tabMurid

	fmt.Scan(&N)
	inputMurid(&daftarMurid, N)

	fmt.Scan(&target)
	idxRes = cariMuridBinary(daftarMurid, N, target)

	if idxRes == -1 {
		fmt.Println("Murid tak terdaftar")
	} else {
		fmt.Printf("Murid terdaftar dan berada pada absen ke-%d\n", idxRes+1)
	}
}

func inputMurid(A *tabMurid, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cariMuridBinary(A tabMurid, n int, x string) int {
	var left, right, mid int
	var found int

	left = 0
	right = n - 1
	found = -1

	for left <= right && found == -1 {

		mid = (left + right) / 2

		if A[mid] == x {
			found = mid
		} else if x < A[mid] {
			right = mid - 1
		} else if x > A[mid] {
			left = mid + 1
		}
	}
	return found
}
