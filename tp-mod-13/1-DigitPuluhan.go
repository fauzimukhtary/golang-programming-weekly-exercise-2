package main

import "fmt"

type arrInt [100]int

func main() {
	var N int
	var arr arrInt

	fmt.Scan(&N)
	inputData(&arr, N)
	fmt.Println("Data sebelum sorting:")
	cetakData(arr, N)
	insSort(&arr, N)
	fmt.Println("Data setelah sorting:")
	cetakData(arr, N)
}

func inputData(A *arrInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func insSort(A *arrInt, n int) {
	var temp, i, j int
	var d1, d2 int

	for i = 1; i < n; i++ {
		j = i
		temp = A[j]
		d1 = temp / 10 % 10
		d2 = A[j-1] / 10 % 10
		for j > 0 && (d1 < d2 || (d1 == d2 && temp < A[j-1])) {
			A[j] = A[j-1]
			j--
			if j > 0 {
				d1 = temp / 10 % 10
				d2 = A[j-1] / 10 % 10
			}
		}
		A[j] = temp
	}
}

func cetakData(A arrInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(A[i])
	}
	fmt.Println()
}
