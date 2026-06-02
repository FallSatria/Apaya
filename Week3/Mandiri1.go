package main

import (
	"fmt"
)
func hitungNilaiAkhir(tugas, uts, uas float64) float64 {
	return (tugas * 0.30) + (uts * 0.30) + (uas * 0.40) 
}

func tentukanNilaiHuruf(nilai float64) string {
	if nilai >= 80 {
		return "A" 
	} else if nilai >= 70 {
		return "B" 
	} else if nilai >= 60 {
		return "C" 
	} else if nilai >= 50 {
		return "D" 
	} else {
		return "E"
	}
}

func main() {
	var nama string
	var tugas, uts, uas float64

	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Nilai Tugas: ")
	fmt.Scanln(&tugas)
	fmt.Print("Nilai UTS: ")
	fmt.Scanln(&uts)
	fmt.Print("Nilai UAS: ")
	fmt.Scanln(&uas)

	nilaiAkhir := hitungNilaiAkhir(tugas, uts, uas)
	nilaiHuruf := tentukanNilaiHuruf(nilaiAkhir)

	fmt.Printf("Nama Mahasiswa: %s\n", nama)
	fmt.Printf("Nilai Akhir: %.0f\n", nilaiAkhir)
	fmt.Printf("Nilai Huruf: %s\n", nilaiHuruf)
}