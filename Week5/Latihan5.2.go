package main

import (
	"fmt"
)

func cetakBintang(sekarang int, n int) {
	if sekarang > n {
		return
	}
	
	for i := 1; i <= sekarang; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	cetakBintang(sekarang + 1, n)
}

func main() {
	var n int
	fmt.Print("Masukan N: ")
	fmt.Scan(&n)
	
	fmt.Println("Keluaran:")
	cetakBintang(1, n)
}