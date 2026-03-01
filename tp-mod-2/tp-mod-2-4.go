package main

import "fmt"

func hitungJamBelajar(n, k int, wBelajar *int) {
	var i, wktMain int

	for i = 1; i <= n; i++ {
		fmt.Scan(&wktMain)
		*wBelajar += k - wktMain
	}

}

func main() {
	var n, k, wBelajar int

	fmt.Println("Input hari yang ingin dihitung:")
	fmt.Scan(&n)
	fmt.Println("Input jam yang dapat digunakan:")
	fmt.Scan(&k)
	fmt.Println("Input waktu main:")

	wBelajar = 0
	hitungJamBelajar(n, k, &wBelajar)

	fmt.Printf("%.1f", float64(wBelajar)/float64(n))
}
