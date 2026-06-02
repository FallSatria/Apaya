package main

import (
	"fmt"
)

type Mahasiswa struct {
	Nama  string
	NIM   string
	Nilai [3]int
}

func main() {
	var mhs Mahasiswa

	fmt.Print("Nama: ")
	fmt.Scanln(&mhs.Nama)
	
	fmt.Print("NIM: ")
	fmt.Scanln(&mhs.NIM)

	for i := 0; i < 3; i++ {
		fmt.Printf("Nilai %d: ", i+1)
		fmt.Scanln(&mhs.Nilai[i])
	}

	totalNilai := mhs.Nilai[0] + mhs.Nilai[1] + mhs.Nilai[2]
	rataRata := float64(totalNilai) / 3.0

	fmt.Printf("Rata-rata Nilai %s (%s) = %.2f\n", mhs.Nama, mhs.NIM, rataRata)
}