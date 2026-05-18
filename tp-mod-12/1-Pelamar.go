package main

import "fmt"

const NMAX int = 100

type Applicant struct {
	id, name     string
	testGrade    float64
	testDuration int
}

type applicants [NMAX]Applicant

func main() {
	var A applicants
	var i, j, n, max int
	var temp Applicant
	var avgDur, avgTest float64

	i = 0
	n = 0

	fmt.Scan(&A[i].id)
	for A[i].id != "END" {
		fmt.Scan(&A[i].name, &A[i].testGrade, &A[i].testDuration)
		avgDur += float64(A[i].testDuration)
		avgTest += A[i].testGrade
		i++
		n++
		fmt.Scan(&A[i].id)
	}

	fmt.Println("Data awal:")
	for i = 0; i < n; i++ {
		fmt.Printf("%-3v %-4v %.1f %-3d\n", A[i].id, A[i].name, A[i].testGrade, A[i].testDuration)
	}

	for i = 0; i < n; i++ {
		max = i

		for j = i + 1; j < n; j++ {
			if A[j].testGrade > A[max].testGrade {
				max = j
			} else if A[j].testGrade == A[max].testGrade {
				if A[j].testDuration < A[max].testDuration {
					max = j
				}
			}
		}

		temp = A[i]
		A[i] = A[max]
		A[max] = temp
	}

	fmt.Println("Data setelah sortir")
	for i = 0; i < n; i++ {
		fmt.Printf("%-3v %-4v %.1f %-3d\n", A[i].id, A[i].name, A[i].testGrade, A[i].testDuration)
	}
	fmt.Printf("Rata-rata durasi tes: %g menit\n", avgDur/float64(n))
	fmt.Printf("Rata-rata nilai tes: %g\n", avgTest/float64(n))
}
