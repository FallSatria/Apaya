package main

import "fmt"

type Buku struct {
	Kode, Judul string
	Harga, Stok int
}

func cariIndeks(buku [5]Buku, kode string) int {
	for i := 0; i < len(buku); i++ {
		if buku[i].Kode == kode {
			return i
		}
	}
	return -1
}

func hitungTotal(harga, jumlah int) int {
	return harga * jumlah
}

func tampilkanBuku(buku [5]Buku) {
	fmt.Println("\n--- Daftar Buku ---")
	for i := 0; i < len(buku); i++ {
		fmt.Printf("%s | %s | Rp%d | Stok: %d\n", buku[i].Kode, buku[i].Judul, buku[i].Harga, buku[i].Stok)
	}
}

func transaksi(buku *[5]Buku) {
	var kode string
	fmt.Print("\nMasukkan Kode Buku: ")
	fmt.Scan(&kode)

	idx := cariIndeks(*buku, kode)

	if idx == -1 {
		fmt.Println("Kode buku tidak ditemukan.")
		return 
	}

	var jumlah int
	fmt.Printf("Buku: %s (Stok: %d)\n", buku[idx].Judul, buku[idx].Stok)
	fmt.Print("Jumlah beli: ")
	fmt.Scan(&jumlah)

	if buku[idx].Stok >= jumlah {
		buku[idx].Stok -= jumlah
		fmt.Printf("Berhasil! Total Bayar: Rp%d\n", hitungTotal(buku[idx].Harga, jumlah))
	} else {
		fmt.Println("Stok tidak mencukupi.")
	}
}

func main() {
	data := [5]Buku{
		{"B1", "Seporsi Mie Ayam Sebelum Meninggal", 50000, 10},
		{"B2", "Filsafat Untuk Orang Pemalas", 40000, 5},
		{"B3", "Psychology Of Money", 60000, 8},
		{"B4", "Calculus For Children", 70000, 3},
		{"B5", "Jujutsu Kaisen Vol 1", 80000, 15},
	}

	tampilkanBuku(data)
	transaksi(&data)

	fmt.Println("\n--- Sisa Stok Terkini ---")
	tampilkanBuku(data)
}