package main

import "fmt"

type Mahasiswa struct {
	NIM  string
	Nama string
}

func main() {
	var n int
	fmt.Scan(&n)

	mhs := make([]Mahasiswa, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mhs[i].NIM, &mhs[i].Nama)
	}

	var targetNIM string
	fmt.Scan(&targetNIM)
	low := 0
	high := n - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if mhs[mid].NIM == targetNIM {
			fmt.Printf("Nama: %s\n", mhs[mid].Nama)
			found = true
			break
		} else if mhs[mid].NIM < targetNIM {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("NIM tidak ditemukan")
	}
}