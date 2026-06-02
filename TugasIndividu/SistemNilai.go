package main

import "fmt"


type Mahasiswa struct {
	NIM        string
	Nama       string
	NilaiTugas float64
	NilaiUTS   float64
	NilaiUAS   float64
}
func hitungNilaiAkhir(tugas, uts, uas float64) float64 {
	return (0.3 * tugas) + (0.3 * uts) + (0.4 * uas)
}
func tentukanGrade(nilaiAkhir float64) string {
	switch {
	case nilaiAkhir >= 85:
		return "A"
	case nilaiAkhir >= 70:
		return "B"
	case nilaiAkhir >= 55:
		return "C"
	case nilaiAkhir >= 40:
		return "D"
	default:
		return "E"
	}
}
func inputMahasiswa(mhs *[3]Mahasiswa) {
	for i := 0; i < len(mhs); i++ {
		fmt.Printf("\nData Mahasiswa ke-%d:\n", i+1)
		fmt.Print("NIM         : ")
		fmt.Scan(&mhs[i].NIM)
		fmt.Print("Nama        : ")
		fmt.Scan(&mhs[i].Nama)
		fmt.Print("Nilai Tugas : ")
		fmt.Scan(&mhs[i].NilaiTugas)
		fmt.Print("Nilai UTS   : ")
		fmt.Scan(&mhs[i].NilaiUTS)
		fmt.Print("Nilai UAS   : ")
		fmt.Scan(&mhs[i].NilaiUAS)
	}
}
// 2. Prosedur: Menampilkan laporan daftar mahasiswa
func tampilkanLaporan(mhs [3]Mahasiswa) {
	fmt.Printf("%-10s %-15s %-15s %-10s\n", "NIM", "Nama", "Nilai Akhir", "Grade")
	for i := 0; i < len(mhs); i++ {
		nilaiAkhir := hitungNilaiAkhir(mhs[i].NilaiTugas, mhs[i].NilaiUTS, mhs[i].NilaiUAS)
		
		grade := tentukanGrade(nilaiAkhir)

		nilaiTanpaDesimal := float64(int(nilaiAkhir))

		fmt.Printf("%-10s %-15s %-15.2f %-10s\n", mhs[i].NIM, mhs[i].Nama, nilaiTanpaDesimal, grade)
	}
}

func main() {
	var dataMahasiswa [3]Mahasiswa
	
	inputMahasiswa(&dataMahasiswa)
	tampilkanLaporan(dataMahasiswa)
}