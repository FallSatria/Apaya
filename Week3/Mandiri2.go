package main

import (
	"fmt"
)

func hitungBiayaKotor(lamaJam int) float64 {
	if lamaJam <= 0 {
		return 0
	} else if lamaJam == 1 {
		return 5000
	} else {
		return 5000 + float64((lamaJam-1)*3000) 
	}
}

func hitungDiskon(biayaKotor float64, lamaJam int) float64 {
	if lamaJam > 5 {
		return biayaKotor * 0.10 
	}
	return 0
}

func main() {
	var kendaraan string
	var lamaParkir int

	fmt.Scan(&kendaraan)
	fmt.Scan(&lamaParkir)

	biayaKotor := hitungBiayaKotor(lamaParkir)
	diskon := hitungDiskon(biayaKotor, lamaParkir)
	
	totalBiayaAkhir := biayaKotor - diskon

	fmt.Printf("Kendaraan: %s\n", kendaraan)
	fmt.Printf("Lama Parkir: %d jam\n", lamaParkir)
	fmt.Printf("Total Biaya Parkir: %.0f\n", totalBiayaAkhir)
}