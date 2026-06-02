package main
import "fmt"

func hitungNilai (nilaiTugas, nilaiUjian float64) int {
	nilaiTugas = nilaiTugas * 0.4
	nilaiUjian = nilaiUjian * 0.6
	nilaiAkhir := nilaiTugas + nilaiUjian
	return int(nilaiAkhir)
}

func cekNilai (nilaiAkhir int) string {
	if nilaiAkhir >= 60 {
		return "LULUS"
	} else {
		return "TIDAK LULUS"
	}
}

func main() {
	var nilaiTugas, nilaiUjian, nilaiAkhir int
	fmt.Scan(&nilaiTugas, &nilaiUjian)
	hasilnilai := hitungNilai(float64(nilaiTugas), float64(nilaiUjian))
	nilaiAkhir = int(hasilnilai)
	hasilcek := cekNilai(nilaiAkhir)
	fmt.Println("Nilai Akhir:", hasilnilai)
	fmt.Println("Status:", hasilcek)
}
