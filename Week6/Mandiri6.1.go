package main

import "fmt"

func cekHarga(pilihan int) int {
	switch pilihan {
	case 1:
		return 5000
	case 2:
		return 7000
	case 3:
		return 2500
	case 4:
		return 4000
	case 5:
		return 2000
	default:
		return 0
	}
}

func main() {
	fmt.Println("1. Susu kotak = Rp 5.000")
	fmt.Println("2. Coklat batang = Rp 7.000")
	fmt.Println("3. Mie instan = Rp 2.500")
	fmt.Println("4. Teh botol = Rp 4.000")
	fmt.Println("5. Kardus bekas = Rp 2.000")
	fmt.Println("0. Berhenti Memilih")


	totalHarga := 0

	for i := 1; i <= 3; i++ {
		var pilihan int
		fmt.Printf("Pilih barang ke-%d: ", i)
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			break
		}

		totalHarga += cekHarga(pilihan)
	}
	fmt.Printf("Total harga: %d\n", totalHarga)
}