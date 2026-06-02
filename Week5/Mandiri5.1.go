package main

import "fmt"


func hitungDigit(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + hitungDigit(n/10)
}

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		bilangan = -bilangan
	}

	hasil := hitungDigit(bilangan)
	
	fmt.Printf("Jumlah digit dari %d adalah %d\n", bilangan, hasil)
}