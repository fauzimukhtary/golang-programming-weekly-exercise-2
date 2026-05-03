package main

import "fmt"

const NMAX int = 50

type Student struct {
	nim, name string
	grade     float64
}

type tabStudents [NMAX]Student

func main() {
	var data tabStudents
	inputDataStudent(&data)
	rapikanData(&data)
	cetakDataBersih(data)
}

func inputDataStudent(t *tabStudents) {
	var i int

	i = 0
	fmt.Scan(&t[i].nim)
	for t[i].nim != "STOP" {
		fmt.Scan(&t[i].name, &t[i].grade)
		i = i + 1
		fmt.Scan(&t[i].nim)
	}
}

func rapikanData(t *tabStudents) {
	var i, j int
	var target string

	i = 0
	for t[i].nim != "STOP" {
		if t[i].nim != "DUPLIKAT" {
			target = t[i].nim
			j = i + 1
			for t[j].nim != "STOP" {
				if t[j].nim == target {
					t[j].nim = "DUPLIKAT"
				}
				j = j + 1
			}
		}
		i = i + 1
	}
}

func cetakDataBersih(t tabStudents) {
	var i int

	i = 0
	fmt.Println("---RESULT---")
	for t[i].nim != "STOP" {
		if t[i].nim != "DUPLIKAT" {
			fmt.Println(t[i].nim, t[i].name, t[i].grade)
		}
		i = i + 1
	}
}
