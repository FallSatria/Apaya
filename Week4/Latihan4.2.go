package main

import (
	"fmt"
)
func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, namaPemenang string
	var maxSoal int = -1
	var minWaktu int = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		var soalSekarang, waktuSekarang int
		hitungSkor(&soalSekarang, &waktuSekarang)

		if soalSekarang > maxSoal || (soalSekarang == maxSoal && waktuSekarang < minWaktu) {
			maxSoal = soalSekarang
			minWaktu = waktuSekarang
			namaPemenang = nama
		}
	}

	if maxSoal != -1 {
		fmt.Printf("%s %d %d\n", namaPemenang, maxSoal, minWaktu)
	}
}