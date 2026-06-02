package main

import "fmt"

type Kandidat struct {
	ID    int
	Suara int
}

func hybridSearch(daftar []Kandidat, targetID int) int {
	kiri := 0
	kanan := len(daftar) - 1
	
	threshold := 4 

	for kiri <= kanan {
		if (kanan - kiri) <= threshold {
			for i := kiri; i <= kanan; i++ {
				if daftar[i].ID == targetID {
					return i
				}
			}
			return -1 
		}

		tengah := (kiri + kanan) / 2

		if daftar[tengah].ID == targetID {
			return tengah
		} else if daftar[tengah].ID < targetID {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	
	return -1
}

func main() {
	var suaraRakyat, suaraSah int
	var nomor int
	var daftarKandidat []Kandidat
	for i := 1; i <= 20; i++ {
		daftarKandidat = append(daftarKandidat, Kandidat{ID: i, Suara: 0})
	}

	for {
		_, err := fmt.Scan(&nomor)
		if err != nil || nomor == 0 {
			break
		}

		suaraRakyat++

		if nomor >= 1 && nomor <= 20 {
			suaraSah++
			indeks := hybridSearch(daftarKandidat, nomor)
			if indeks != -1 {
				daftarKandidat[indeks].Suara++
			}
		}
	}
	fmt.Printf("Suara rakyat: %d\n", suaraRakyat)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for _, k := range daftarKandidat {
		if k.Suara > 0 {
			fmt.Printf("%d: %d\n", k.ID, k.Suara)
		}
	}
}