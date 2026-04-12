package main

import "fmt"

func main() {
	var konstanta int
	var res int
	fmt.Scan(&konstanta)
	res = smallestRepunitDivByK(konstanta)
	fmt.Println(res)
}

func findRepunit(k, length, rem int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	} else if rem%k == 0 {
		return length
	} else {
		return findRepunit(k, length+1, (rem*10+1)%k)
	}
}

func smallestRepunitDivByK(k int) int {
	return findRepunit(k, 1, 1)
}
