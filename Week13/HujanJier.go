package main

import "fmt"

func main() {
	var hujan []int
	var input int

	for {
		fmt.Scan(&input)
		if input == -999 {
			break
		}
		hujan = append(hujan, input)
	}
	if len(hujan) == 0 {
		return
	}
	jumlahHari := len(hujan)
	sum := 0
	max := hujan[0]
	maxHari := 1
	min := hujan[0]
	minHari := 1
	peningkatan := 0
	for i, val := range hujan {
		sum += val
		if val > max {
			max = val
			maxHari = i + 1
		}
		if val < min {
			min = val
			minHari = i + 1
		}
		if i > 0 && val > hujan[i-1] {
			peningkatan++
		}
	}
	rataRata := float64(sum) / float64(jumlahHari)
	diAtasRata := 0
	for _, val := range hujan {
		if float64(val) > rataRata {
			diAtasRata++
		}
	}
	fmt.Printf("Jumlah hari: %d\n", jumlahHari)
	fmt.Printf("Rata-rata hujan: %.2f\n", rataRata)
	fmt.Printf("Curah hujan tertinggi: %d (hari ke-%d)\n", max, maxHari)
	fmt.Printf("Curah hujan terendah: %d (hari ke-%d)\n", min, minHari)
	fmt.Printf("Hari di atas rata-rata: %d\n", diAtasRata)
	fmt.Printf("Jumlah peningkatan: %d\n", peningkatan)
}