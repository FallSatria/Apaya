package main

import "fmt"

const NMAX = 1000000
var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
	indeks := posisi(n, k)
	if indeks == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(indeks)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i
		}
	}

	return -1 
}