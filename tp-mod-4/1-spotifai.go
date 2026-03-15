package main

import "fmt"

const NMAX int = 999

type song struct {
	judul, penyanyi          string
	durasiMenit, durasiDetik int
}

type tabLagu struct {
	totalDurasi int
	arrLagu     [NMAX]song
}

func main() {
	var n int
	var arrayLagu tabLagu
	fmt.Scan(&n)
	inputLagu(&arrayLagu, n)
	printLagu(arrayLagu, n)
}

func inputLagu(s *tabLagu, n int) {
	var i int
	s.totalDurasi = 0
	for i = 0; i < n; i++ {
		fmt.Print("\nJudul Lagu: ")
		fmt.Scan(&s.arrLagu[i].judul)
		fmt.Print("Nama Penyanyi: ")
		fmt.Scan(&s.arrLagu[i].penyanyi)
		fmt.Print("Durasi Lagu (menit detik): ")
		fmt.Scan(&s.arrLagu[i].durasiMenit, &s.arrLagu[i].durasiDetik)
		s.totalDurasi += (s.arrLagu[i].durasiMenit*60 + s.arrLagu[i].durasiDetik)
	}
}

func printLagu(s tabLagu, n int) {
	var i int
	fmt.Println()
	fmt.Print("+---------------+---------------+----------+\n")
	fmt.Print("| Judul Lagu    | Penyanyi      | Durasi   |\n")
	fmt.Print("+---------------+---------------+----------+\n")
	for i = 0; i < n; i++ {
		fmt.Printf("| %v\t", s.arrLagu[i].judul)
		if len(s.arrLagu[i].judul) < 6 {
			fmt.Print("\t")
		}

		fmt.Printf("| %v\t", s.arrLagu[i].penyanyi)
		if len(s.arrLagu[i].penyanyi) < 6 {
			fmt.Print("\t")
		}

		fmt.Printf("| %02d", s.arrLagu[i].durasiMenit)
		fmt.Printf(":%02d    |\n", s.arrLagu[i].durasiDetik)
	}
	fmt.Print("+---------------+---------------+----------+\n")
	fmt.Print("| Total Durasi Lagu\t\t")
	fmt.Printf("| %02d", s.totalDurasi/3600)
	fmt.Printf(":%02d", s.totalDurasi%3600/60)
	fmt.Printf(":%02d |\n", s.totalDurasi%60)
	fmt.Print("+---------------+---------------+----------+\n")
	fmt.Println()
}
