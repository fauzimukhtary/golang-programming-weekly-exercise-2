package main

import "fmt"

const MAX int = 99

type arrString [MAX]string

func main() {
	var N int
	var arr arrString

	fmt.Scan(&N)
	bacaData(&arr, N)
	selectionSortAscend(&arr, N)
	selectionSortDescend(&arr, N)
}

func bacaData(A *arrString, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A arrString, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("%v", A[i])
		if i != n-1 {
			fmt.Print(", ")
		} else {
			fmt.Print(".\n")
		}
	}
}

func selectionSortAscend(A *arrString, n int) {
	var i, j, min int
	var temp string

	for i = 0; i < n; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}

	fmt.Println("Data setelah diurutkan secara Ascending:")
	cetakData(*A, n)
}

func selectionSortDescend(A *arrString, n int) {
	var i, j, max int
	var temp string

	for i = 0; i < n; i++ {
		max = i
		for j = i + 1; j < n; j++ {
			if A[j] > A[max] {
				max = j
			}
		}
		temp = A[i]
		A[i] = A[max]
		A[max] = temp
	}

	fmt.Println("Data setelah diurutkan secara Descending:")
	cetakData(*A, n)
}
