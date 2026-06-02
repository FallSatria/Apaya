package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	products := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&products[i])
	}

	var target string
	fmt.Scan(&target)

	found := false
	for i := 0; i < n; i++ {
		if products[i] == target {
			fmt.Printf("Produk ditemukan di indeks %d\n", i)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Produk tidak ditemukan")
	}
}