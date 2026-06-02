package main

import "fmt"

func hitungLuas(pilihan int) {
	switch pilihan {
	case 1:
		var sisi float64
		fmt.Print("Sisi: ")
		fmt.Scanln(&sisi)
		fmt.Printf("Luas bangun datar persegi: %.0f\n", sisi*sisi)

	case 2:
		var panjang, lebar float64
		fmt.Print("Panjang: ")
		fmt.Scanln(&panjang)
		fmt.Print("Lebar: ")
		fmt.Scanln(&lebar)
		fmt.Printf("Luas bangun datar persegi panjang: %.0f\n", panjang*lebar)

	case 3:
		var alas, tinggi float64
		fmt.Print("Alas: ")
		fmt.Scanln(&alas)
		fmt.Print("Tinggi: ")
		fmt.Scanln(&tinggi)
		fmt.Printf("Luas bangun datar segitiga: %.1f\n", 0.5*alas*tinggi)

	default:
		fmt.Println("Pilihan bangun datar tidak valid.")
	}
}

func main() {
	var pilihan int

	fmt.Print("Pilih bangun datar (1/2/3): ")
	fmt.Scanln(&pilihan)

	hitungLuas(pilihan)
}