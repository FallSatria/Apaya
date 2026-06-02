package main

import "fmt"

func main() {
	var suaraMasuk int = 0
	var suaraSah int = 0
	var votes [21]int 

	var nomor int
	for {
		_, err := fmt.Scan(&nomor)
		if err != nil || nomor == 0 {
			break
		}

		suaraMasuk++
		if nomor >= 1 && nomor <= 20 {
			suaraSah++
			votes[nomor]++
		}
	}
	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	for i := 1; i <= 20; i++ {
		if votes[i] > 0 {
			fmt.Printf("%d: %d\n", i, votes[i])
		}
	}
}