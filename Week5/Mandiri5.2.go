package main

import "fmt"

func segitigaTerbalik(n int) {
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()

	segitigaTerbalik(n - 1)
}

func main() {
	var n int

	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	segitigaTerbalik(n)
}