package main

import (
	"fmt"
)

type DataMahasiswa struct {
	Nama     string
	Nilai    [3]int
	RataRata float64
}

func main() {
	var jumlah int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scanln(&jumlah)

	if jumlah > 10 {
		fmt.Println("Jumlah mahasiswa maksimal adalah 10.")
		return
	}

	var daftarMhs [10]DataMahasiswa
	var maxRata float64 = -1.0
	var namaTertinggi string

	for i := 0; i < jumlah; i++ {
		fmt.Printf("--- Mahasiswa ke-%d ---\n", i+1)
		
		fmt.Print("Masukkan nama mahasiswa: ")
		fmt.Scanln(&daftarMhs[i].Nama)

		total := 0
		for j := 0; j < 3; j++ {
			fmt.Printf("Masukkan nilai tugas %d: ", j+1)
			fmt.Scanln(&daftarMhs[i].Nilai[j])
			total += daftarMhs[i].Nilai[j]
		}

		daftarMhs[i].RataRata = float64(total) / 3.0

		if daftarMhs[i].RataRata > maxRata {
			maxRata = daftarMhs[i].RataRata
			namaTertinggi = daftarMhs[i].Nama
		}
	}

	fmt.Println("\nDATA NILAI MAHASISWA")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%d. %s Rata-rata: %.2f\n", i+1, daftarMhs[i].Nama, daftarMhs[i].RataRata)
	}
	
	fmt.Printf("Nilai tertinggi: %s (%.2f)\n", namaTertinggi, maxRata)
}