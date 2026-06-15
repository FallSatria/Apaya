package main

import "fmt"


func main() {
	var skor []int
	var input int

	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		skor = append(skor, input)
	}
	if len(skor) == 0 {
		return
	}
	jumlahSesi := len(skor)
	max := skor[0]
	maxSesi := 1
	min := skor[0]
	minSesi := 1
	peningkatan := 0
	penurunan := 0
	streakTerlama := 0
	streakSekarang := 0
	for i := 0; i < jumlahSesi; i++ {
		if skor[i] > max {
			max = skor[i]
			maxSesi = i + 1
		}
		if skor[i] < min {
			min = skor[i]
			minSesi = i + 1
		}
		if i == 0 {
			continue
		}
		if skor[i] > skor[i-1] {
			peningkatan++
			streakSekarang++
			if streakSekarang > streakTerlama {
				streakTerlama = streakSekarang
			}
		} else if skor[i] < skor[i-1] {
			penurunan++
			streakSekarang = 0 
		} else {
			streakSekarang = 0 
		}
	}

	// Cetak Keluaran
	fmt.Printf("Jumlah sesi: %d\n", jumlahSesi)
	fmt.Printf("Nilai tertinggi: %d (sesi ke-%d)\n", max, maxSesi)
	fmt.Printf("Nilai terendah: %d (sesi ke-%d)\n", min, minSesi)
	fmt.Printf("Jumlah peningkatan: %d\n", peningkatan)
	fmt.Printf("Jumlah penurunan: %d\n", penurunan)
	fmt.Printf("Streak peningkatan terlama: %d\n", streakTerlama)
}