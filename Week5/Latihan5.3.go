package main

import "fmt"

func cetakFaktor(n int, pembagi int) {

	if pembagi > n {
		fmt.Println() 
		return
	}
	
	if n%pembagi == 0 {
		fmt.Printf("%d ", pembagi)
	}
	cetakFaktor(n, pembagi+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	fmt.Print("Keluaran: ")
	cetakFaktor(n, 1)
}