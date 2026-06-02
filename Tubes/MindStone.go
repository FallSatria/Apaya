package main 
import "fmt"

const NMAX = 100
var data [NMAX]int

type pekerja struct {
	namaPekerja string
	namaTugas string
	mood string
	progress string
	skorStress int
}


func TambahPencapaian(arr *[NMAX]pekerja, jumlah *int) {
	if *jumlah >= NMAX {
		fmt.Println("Data sudah penuh")
		return 
	}else if *jumlah < NMAX {
		fmt.Scan(&arr[*jumlah].namaPekerja, &arr[*jumlah].namaTugas, &arr[*jumlah].mood, &arr[*jumlah].progress, &arr[*jumlah].skorStress)
		*jumlah++
	}
}

func OutputPencapaian(arr *[NMAX]pekerja, jumlah *int) {
	if *jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		return 
	}
	for i := 0; i < *jumlah; i++ {
		fmt.Printf("Nama Pekerja: %s,", arr[i].namaPekerja)
		fmt.Printf("Nama Tugas: %s,", arr[i].namaTugas)
		fmt.Printf("Mood: %s,", arr[i].mood)
		fmt.Printf("Progress: %s,", arr[i].progress)
		fmt.Printf("Skor Stress: %d\n", arr[i].skorStress)
	}
}

func UpdatePencapaian(arr *[NMAX]pekerja, jumlah *int) {
	if *jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		return 
	}
	var nomor int
	fmt.Print("Masukkan nomor pencapaian yang ingin diupdate: ")
	fmt.Scan(&nomor)
	if nomor < 1 || nomor > *jumlah {
		fmt.Println("Nomor pencapaian tidak valid")
		return 
	}
	fmt.Printf("Masukkan data baru untuk pencapaian nomor %d:\n", nomor)
	fmt.Scan(&arr[nomor-1].namaPekerja, &arr[nomor-1].namaTugas, &arr[nomor-1].mood, &arr[nomor-1].progress, &arr[nomor-1].skorStress)
}

func main() {
	var pilihan int
	var data [NMAX]pekerja
	var jumlahData int = 0 

	for {
		fmt.Println("==== Aplikasi MindStone ====")
		fmt.Println("1. Tambah Pencapaian")
		fmt.Println("2. Lihat Pencapaian")
		fmt.Println("3. Update Pencapaian")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		switch pilihan {
			case 1:
				TambahPencapaian(&data, &jumlahData)
			case 0:
				return
			default:
				fmt.Println("Pilihan tidak valid, silakan coba lagi.")

		}
	}

}
