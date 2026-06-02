package main

import "fmt"


// Konstanta batas maksimum array sesuai spesifikasi soal
const nMAX = 1000

// Deklarasi tipe bentukan Lulusan
type Lulusan struct {
	nama     string
	nim      string
	semester int
	eprt     float64
	ipk      float64
}

// Deklarasi tipe array tabLulus
type TabLulus [nMAX]Lulusan

// 1. Procedure untuk pengisian array wisudawan
// Menggunakan pointer (*TabLulus dan *int) agar perubahan data tersimpan ke variabel asli
func isiData(t *TabLulus, n *int) {
	*n = 0
	var nim string

	fmt.Print("Masukkan NIM (ketik 'none' untuk berhenti): ")
	fmt.Scan(&nim)

	for nim != "none" && *n < nMAX {
		t[*n].nim = nim

		fmt.Print("Masukkan Nama, EPRT, Semester, dan IPK (pisahkan dengan spasi): ")
		fmt.Scan(&t[*n].nama, &t[*n].eprt, &t[*n].semester, &t[*n].ipk)

		*n++ // Increment jumlah mahasiswa

		fmt.Print("Masukkan NIM (ketik 'none' untuk berhenti): ")
		fmt.Scan(&nim)
	}
}

// 2a. Function untuk mencari EPRT tertinggi
func maxEPRT(t TabLulus, n int) float64 {
	if n == 0 {
		return 0 // Antisipasi jika array kosong
	}
	
	max := t[0].eprt // Jadikan data pertama sebagai acuan nilai ekstrim
	
	for i := 1; i < n; i++ {
		// Jika nilai acuan lebih kecil dari data saat ini, update nilai max
		if max < t[i].eprt {
			max = t[i].eprt
		}
	}
	return max
}

// 2b. Function untuk mencari IPK terendah
func minIPK(t TabLulus, n int) float64 {
	if n == 0 {
		return 0
	}
	
	min := t[0] // Menyimpan seluruh record data mahasiswa pertama sebagai acuan
	
	for i := 1; i < n; i++ {
		// Catatan perbaikan: Pada pseudocode slide 12 tertulis min.ipk < t[i].ipk yang akan mencari nilai Max.
		// Untuk mencari nilai minimum, kondisinya harus dibalik menjadi:
		if t[i].ipk < min.ipk {
			min = t[i]
		}
	}
	return min.ipk
}

// 2c. Function untuk mencari rata-rata semester lulusan
func rataan(t TabLulus, n int) float64 {
	if n == 0 {
		return 0
	}
	
	var jum float64 = 0
	
	// Iterasi dari indeks 0 hingga n-1
	for i := 0; i < n; i++ {
		jum += float64(t[i].semester)
	}
	return jum / float64(n)
}

// 3. Program Utama
func main() {
	var dataMHS TabLulus
	var nMHS int

	
	fmt.Println("--- Pengisian Data Lulusan ---")
	isiData(&dataMHS, &nMHS)

	if nMHS > 0 {
		
		eprtTertinggi := maxEPRT(dataMHS, nMHS)
		ipkTerendah := minIPK(dataMHS, nMHS)
		rerataSemester := rataan(dataMHS, nMHS)

		fmt.Println("\n--- Hasil Laporan Wisudawan ---")
		fmt.Printf("EPRT Tertinggi       : %.2f\n", eprtTertinggi)
		fmt.Printf("IPK Terendah         : %.2f\n", ipkTerendah)
		fmt.Printf("Rata-rata Semester   : %.2f\n", rerataSemester)
	} else {
		fmt.Println("Tidak ada data wisudawan yang dimasukkan.")
	}
}