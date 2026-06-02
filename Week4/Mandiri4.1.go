package main

import "fmt"

func hitungDiskon(totalBelanja float64) (int, float64) {
	var diskon int

	if totalBelanja >= 500000 {
		diskon = 20
	} else if totalBelanja >= 300000 {
		diskon = 15
	} else if totalBelanja >= 100000 {
		diskon = 10
	} else {
		diskon = 0
	}

	potongan := totalBelanja * (float64(diskon) / 100)
	totalBayar := totalBelanja - potongan

	return diskon, totalBayar
}

func main() {
	var nama string
	var totalBelanja float64

	fmt.Print("Nama pelanggan: ")
	fmt.Scanln(&nama)
	fmt.Print("Total belanja: ")
	fmt.Scanln(&totalBelanja)

	persenDiskon, totalBayar := hitungDiskon(totalBelanja)

	fmt.Printf("\nPelanggan: %s\n", nama)
	fmt.Printf("Diskon: %d%%\n", persenDiskon)
	fmt.Printf("Total bayar: %.0f\n", totalBayar)
}