package main

import "fmt"
type Peserta struct {
	Nama          string
	NIM            string
	NilaiTeori     int
	NilaiPraktikum int
}

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		return
	}

	peserta := make([]Peserta, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&peserta[i].Nama, &peserta[i].NIM, &peserta[i].NilaiTeori, &peserta[i].NilaiPraktikum)
	}
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if peserta[j].NilaiTeori < peserta[minIdx].NilaiTeori {
				minIdx = j
			} else if peserta[j].NilaiTeori == peserta[minIdx].NilaiTeori {
				if peserta[j].NIM < peserta[minIdx].NIM {
					minIdx = j
				}
			}
		}
		peserta[i], peserta[minIdx] = peserta[minIdx], peserta[i]
	}
	kuota := n / 2
	lolos := peserta[n-kuota:]
	for i := 1; i < len(lolos); i++ {
		key := lolos[i]
		j := i - 1
		for j >= 0 {
			if lolos[j].NilaiPraktikum < key.NilaiPraktikum {
				lolos[j+1] = lolos[j]
				j--
			} else if lolos[j].NilaiPraktikum == key.NilaiPraktikum && lolos[j].Nama < key.Nama {
				lolos[j+1] = lolos[j]
				j--
			} else {
				break
			}
		}
		lolos[j+1] = key
	}
	fmt.Println("Hasil Seleksi Akhir ---")
	for _, p := range lolos {
		fmt.Printf("%s - %s [Teori: %d, Praktikum: %d]\n", p.Nama, p.NIM, p.NilaiTeori, p.NilaiPraktikum)
	}
}