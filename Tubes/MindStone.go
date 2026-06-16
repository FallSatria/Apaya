package main
import "fmt"

const NMAX = 100

var data [NMAX]int

type waktu struct {
	hari  int
	bulan int
	tahun int
}

type pekerja struct {
	namaPekerja string
	namaTugas   string
	mood        int
	progress    int
	skorStress  int
	deadline    waktu
	catatan     string
}

func TambahPencapaian(arr *[NMAX]pekerja, jumlah *int) {
	if *jumlah >= NMAX {
		fmt.Println("Data sudah penuh")
		fmt.Println()
		return
	} else if *jumlah < NMAX {
		fmt.Println("Masukkan data pencapaian baru:")
		fmt.Print("nama pekerja		: ")
		fmt.Scan(&arr[*jumlah].namaPekerja)
		fmt.Print("nama tugas		: ")
		fmt.Scan(&arr[*jumlah].namaTugas)
		fmt.Print("mood bagus-jelek(1-10)	: ")
		fmt.Scan(&arr[*jumlah].mood)		
		fmt.Print("progress		: ")
		fmt.Scan(&arr[*jumlah].progress)
		if progress := arr[*jumlah].progress; progress == 100 {
			fmt.Print("catatan			: ")
			fmt.Scan(&arr[*jumlah].catatan)
		}
		fmt.Print("skor stress		: ")
		fmt.Scan(&arr[*jumlah].skorStress)
		fmt.Print("hari deadline		: ")
		fmt.Scan(&arr[*jumlah].deadline.hari)
		fmt.Print("bulan deadline		: ")
		fmt.Scan(&arr[*jumlah].deadline.bulan)
		fmt.Print("tahun deadline		: ")
		fmt.Scan(&arr[*jumlah].deadline.tahun)
		fmt.Println()
		*jumlah++
	}
}
func Moodtostring(mood int) string {
	switch mood {
	case 1:
		return "sangat bagus"
	case 2, 3:
		return "bagus"
	case 4:
		return "cukup bagus"
	case 5:
		return "biasa saja"
	case 6, 7:
		return "cukup jelek"
	case 8, 9:
		return "jelek"
	case 10:
		return "sangat jelek"
	default:
		return "Progres belum selesai atau mood tidak tersedia"
	}
}

func OutputPencapaian(arr *[NMAX]pekerja, jumlah *int) {
	if *jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		fmt.Println()
		return
	}
	for i := 0; i < *jumlah; i++ {
		fmt.Printf("Pencapaian #%d\n", i+1)
		fmt.Printf("Nama Pekerja	: %s\n", arr[i].namaPekerja)
		fmt.Printf("Nama Tugas	: %s\n", arr[i].namaTugas)
		fmt.Printf("Mood		: %s\n", Moodtostring(arr[i].mood))
		fmt.Printf("Progress	: %d%%\n", arr[i].progress)
		if arr[i].progress == 100 {
			fmt.Printf("Catatan		: %s\n", arr[i].catatan)
		}
		fmt.Printf("Skor Stress	: %d\n", arr[i].skorStress)
		fmt.Printf("Deadline	: %02d/%02d/%04d\n", arr[i].deadline.hari, arr[i].deadline.bulan, arr[i].deadline.tahun)
		fmt.Println()
	}
}

func UpdatePencapaian(arr *[NMAX]pekerja, jumlah *int) {
	if *jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		fmt.Println()
		return
	}
	var nomor int
	fmt.Print("Masukkan nomor pencapaian yang ingin diupdate: ")
	fmt.Println()
	fmt.Scan(&nomor)
	if nomor < 1 || nomor > *jumlah {
		fmt.Println("Nomor pencapaian tidak valid")
		fmt.Println()
		return
	}
	fmt.Printf("Masukkan data baru untuk pencapaian nomor %d:\n", nomor)
	fmt.Print("Nama Pekerja		: ")
	fmt.Scan(&arr[nomor-1].namaPekerja)
	fmt.Print("Nama Tugas		: ")
	fmt.Scan(&arr[nomor-1].namaTugas)
	fmt.Print("Mood bagus-jelek(1-10)	: ")
	fmt.Scan(&arr[nomor-1].mood)
	fmt.Print("Progress		: ")
	fmt.Scan(&arr[nomor-1].progress)
	if progress := arr[nomor-1].progress; progress == 100 {
		fmt.Print("catatan			: ")
		fmt.Scan(&arr[nomor-1].catatan)
	}
	fmt.Print("Skor Stress		: ")
	fmt.Scan(&arr[nomor-1].skorStress)
	fmt.Print("Hari Deadline		: ")
	fmt.Scan(&arr[nomor-1].deadline.hari)
	fmt.Print("Bulan Deadline		: ")
	fmt.Scan(&arr[nomor-1].deadline.bulan)
	fmt.Print("Tahun Deadline		: ")
	fmt.Scan(&arr[nomor-1].deadline.tahun)
	fmt.Println()
}

func HapusPencapaian(arr *[NMAX]pekerja, jumlah *int) {
	if *jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")

		return
	}
	var nomor int
	fmt.Print("Masukkan nomor pencapaian yang ingin dihapus: ")
	fmt.Scan(&nomor)
	fmt.Println()
	if nomor < 1 || nomor > *jumlah {
		fmt.Println("Nomor pencapaian tidak valid")
		fmt.Println()
		return
	}
	fmt.Print("Masukkan nomor pencapaian yang ingin dihapus: ")
	fmt.Scan(&nomor)
	fmt.Println()
	if nomor < 1 || nomor > *jumlah {
		fmt.Println("Nomor pencapaian tidak valid")
		fmt.Println()
		return
	}
	indexhapus := nomor - 1
	for i := indexhapus; i < *jumlah-1; i++ {
		arr[i] = arr[i+1]
	}
	*jumlah--
}

func UrutkanMoodSelection(arr *[NMAX]pekerja, jumlah int) {
	if jumlah <= 1 {
		fmt.Println("Data tidak cukup untuk diurutkan")
		fmt.Println()
		return
	}
	salinan := *arr

	for i := 0; i < jumlah-1; i++ {
		idxMax := i
		for j := i + 1; j < jumlah; j++ {
			if salinan[j].mood < salinan[idxMax].mood {
				idxMax = j
			}
		}
		if idxMax != i {
			salinan[i], salinan[idxMax] = salinan[idxMax], salinan[i]
		}
	}
	fmt.Println("HASIL URUTAN BERDASARKAN MOOD PALING TINGGI")
	fmt.Println()
	OutputPencapaian(&salinan, &jumlah)
}

func UrutkanMoodInsertion(arr *[NMAX]pekerja, jumlah int) {
	if jumlah <= 1 {
		fmt.Println("Data tidak cukup untuk diurutkan")
		fmt.Println()
		return
	}
	salinan := *arr
	for i := 1; i < jumlah; i++ {
		key := salinan[i]
		j := i - 1
		for j >= 0 && salinan[j].mood < key.mood {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = key
	}
	fmt.Println("HASIL URUTAN BERDASARKAN MOOD PALING RENDAH")
	fmt.Println()
	OutputPencapaian(&salinan, &jumlah)
}

func DuluanGaNih(A waktu, B waktu) bool {
	if A.tahun != B.tahun {
		return A.tahun < B.tahun
	}
	if A.bulan != B.bulan {
		return A.bulan < B.bulan
	}
	return A.hari < B.hari
}

func UrutkanDeadlineTerdekat(arr *[NMAX]pekerja, jumlah int) {
	if jumlah <= 1 {
		fmt.Println("Data tidak cukup untuk diurutkan")
		fmt.Println()
		return
	}

	salinan := *arr

	for i := 1; i < jumlah; i++ {
		key := salinan[i]
		j := i - 1
		for j >= 0 && DuluanGaNih(key.deadline, salinan[j].deadline) {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = key
	}
	fmt.Println("HASIL URUTAN BERDASARKAN DEADLINE PALING DEKAT")
	fmt.Println()
	OutputPencapaian(&salinan, &jumlah)
}

func BelakanganAjaLah(A waktu, B waktu) bool {
	if A.tahun != B.tahun {
		return A.tahun > B.tahun
	}
	if A.bulan != B.bulan {
		return A.bulan > B.bulan
	}
	return A.hari > B.hari
}

func UrutkanDeadlineTerjauh(arr *[NMAX]pekerja, jumlah int) {
	if jumlah <= 1 {
		fmt.Println("Data tidak cukup untuk diurutkan")
		fmt.Println()
		return
	}

	salinan := *arr
	for i := 0; i < jumlah-1; i++ {
		idxMin := i
		for j := i + 1; j < jumlah; j++ {
			if BelakanganAjaLah(salinan[j].deadline, salinan[idxMin].deadline) {
				idxMin = j
			}
		}
		if idxMin != i {
			salinan[i], salinan[idxMin] = salinan[idxMin], salinan[i]
		}
	}
	fmt.Println("HASIL URUTAN BERDASARKAN DEADLINE PALING JAUH")
	fmt.Println()
	OutputPencapaian(&salinan, &jumlah)
}

func CariBerdasarkanTugas(arr *[NMAX]pekerja, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		fmt.Println()
		return
	}
	var keyword string
	fmt.Print("Masukkan nama tugas yang ingin dicari: ")
	fmt.Scan(&keyword)
	fmt.Println()
	found := false
	for i := 0; i < jumlah; i++ {
		if arr[i].namaTugas == keyword {
			fmt.Println("Pencapaian ditemukan:")
			fmt.Println("nama		:", arr[i].namaPekerja)
			fmt.Println("tugas		:", arr[i].namaTugas)
			fmt.Println("mood		:", Moodtostring(arr[i].mood))
			fmt.Println("progress	:", arr[i].progress, "%")
				if arr[i].progress == 100 {
					fmt.Printf("catatan			: %s\n", arr[i].catatan)
				}
			fmt.Println("skor stress	:", arr[i].skorStress)
			fmt.Printf("deadline	: %02d/%02d/%04d\n", arr[i].deadline.hari, arr[i].deadline.bulan, arr[i].deadline.tahun)
			found = true
		}
	}
	if !found {
		fmt.Println("Tugas tidak ditemukan")
	}
	fmt.Println()
}

func CariBerdasarkanDeadline(arr *[NMAX]pekerja, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		fmt.Println()
		return
	}
	var hari, bulan, tahun int
	fmt.Println("Masukkan tanggal deadline yang ingin dicari:  ")
	fmt.Print("Hari		: ")
	fmt.Scan(&hari)
	fmt.Print("Bulan		: ")
	fmt.Scan(&bulan)
	fmt.Print("Tahun		: ")
	fmt.Scan(&tahun)
	fmt.Println()
	found := false
	for i := 0; i < jumlah; i++ {
		if arr[i].deadline.hari == hari && arr[i].deadline.bulan == bulan && arr[i].deadline.tahun == tahun {
			fmt.Println("Pencapaian ditemukan:")
			fmt.Println("nama		:", arr[i].namaPekerja)
			fmt.Println("tugas		:", arr[i].namaTugas)
			fmt.Println("mood		:", Moodtostring(arr[i].mood))
			fmt.Println("progress	:", arr[i].progress, "%")
			if arr[i].progress == 100 {
				fmt.Printf("catatan			: %s\n", arr[i].catatan)
			}
			fmt.Println("skor stress	:", arr[i].skorStress)
			fmt.Printf("deadline	: %02d/%02d/%04d\n", arr[i].deadline.hari, arr[i].deadline.bulan, arr[i].deadline.tahun)
			fmt.Println()
			found = true
		}
	}
	if !found {
		fmt.Println("Deadline tidak ditemukan")
	}
	fmt.Println()
}

func CariBerdasarkanTugasBinary(arr *[NMAX]pekerja, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		return
	}
	var keyword string
	fmt.Println("Masukan tugas yang ingin dicari: ")
	fmt.Scan(&keyword)
	salinan := *arr
	for i := 1; i < jumlah; i++ {
		key := salinan[i]
		j := i - 1
		for j >= 0 && salinan[j].namaTugas > key.namaTugas {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = key
	}
	kiri := 0
	kanan := jumlah - 1
	found := false

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if salinan[tengah].namaTugas == keyword {
			fmt.Println("\n===Pencapaian Ditemukan===")
			fmt.Println("Nama Pekerja	:", salinan[tengah].namaPekerja)
			fmt.Println("Nama Tugas	:", salinan[tengah].namaTugas)
			fmt.Println("Mood		:", Moodtostring(salinan[tengah].mood))
			fmt.Printf("Progress	: %d%%\n", salinan[tengah].progress)
			if salinan[tengah].progress == 100 {
				fmt.Printf("Catatan			: %s\n", salinan[tengah].catatan)
			}
			fmt.Printf("Skor Stress	: %d\n", salinan[tengah].skorStress)
			fmt.Printf("Deadline	: %02d/%02d/%04d\n", salinan[tengah].deadline.hari, salinan[tengah].deadline.bulan, salinan[tengah].deadline.tahun)
			found = true
			break
		} else if salinan[tengah].namaTugas < keyword {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}

	}
	if !found {
		fmt.Println("\nTugas tidak ditemukan")
	}
}

func CariBerdasarkanDeadlineBinary(arr *[NMAX]pekerja, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		return
	}
	var hari, bulan, tahun int
	fmt.Println("Masukkan tanggal deadline yang ingin dicari:  ")
	fmt.Print("Hari		: ")
	fmt.Scan(&hari)
	fmt.Print("Bulan		: ")
	fmt.Scan(&bulan)
	fmt.Print("Tahun		: ")
	fmt.Scan(&tahun)
	targetWaktu := waktu{hari, bulan, tahun}
	salinan := *arr
	for i := 1; i < jumlah; i++ {
		key := salinan[i]
		j := i - 1
		for j >= 0 && DuluanGaNih(key.deadline, salinan[j].deadline) {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = key
	}
	kiri := 0
	kanan := jumlah - 1
	found := false
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		waktuTengah := salinan[tengah].deadline
		if salinan[tengah].deadline.hari == hari && salinan[tengah].deadline.bulan == bulan && salinan[tengah].deadline.tahun == tahun {
			fmt.Println("\n===Pencapaian Ditemukan===")
			fmt.Println("Nama Pekerja	:", salinan[tengah].namaPekerja)
			fmt.Println("Nama Tugas	:", salinan[tengah].namaTugas)
			fmt.Println("Mood		:", Moodtostring(salinan[tengah].mood))
			fmt.Printf("Progress	: %d%%\n", salinan[tengah].progress)
			if salinan[tengah].progress == 100 {
				fmt.Printf("Catatan			: %s\n", salinan[tengah].catatan)
			}
			fmt.Printf("Skor Stress	: %d\n", salinan[tengah].skorStress)
			fmt.Printf("Deadline	: %02d/%02d/%04d\n", salinan[tengah].deadline.hari, salinan[tengah].deadline.bulan, salinan[tengah].deadline.tahun)
			found = true
			break
		} else if DuluanGaNih(targetWaktu, waktuTengah) {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if !found {
		fmt.Println("\nDeadline tidak ditemukan")
	}

}

func RataRataStressMingguan(arr *[NMAX]pekerja, jumlah int) float64 {
	if jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		return 0.0
	}
	totalStress := 0
	minhari := arr[0].deadline.tahun*365 + arr[0].deadline.bulan*30 + arr[0].deadline.hari
	maxhari := minhari
	for i := 0; i < jumlah; i++ {
		totalStress += arr[i].skorStress
		carrhari := arr[i].deadline.tahun*365 + arr[i].deadline.bulan*30 + arr[i].deadline.hari
		if carrhari < minhari {
			minhari = carrhari
		}
		if carrhari > maxhari {
			maxhari = carrhari
		}
	}
	selisihHari := maxhari - minhari
	jumlahMinggu := float64(selisihHari) / 7.0
	rataRata := float64(totalStress) / float64(jumlahMinggu)
	return rataRata
}

func TampilkanStatistikPencapaian(arr *[NMAX]pekerja, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data pencapaian")
		return 
	}
	totalStress := 0
	jumlahSelesai := 0

	fmt.Println("===LAPORAN & STATISTIK MINDSTONE===")
	fmt.Println("Daftar Tugas yang Sudah selesai:")
	for i := 0; i < jumlah; i++ {
		totalStress += arr[i].skorStress
		if arr[i].progress == 100 {
			jumlahSelesai++
			fmt.Printf("- Tugas : %s (Dikerjakan oleh %s)\n", arr[i].namaTugas, arr[i].namaPekerja)
		}
	}

	if jumlahSelesai == 0 {
		fmt.Println("Belum ada tugas yang selesai")
	}
	rataRataStress := float64(totalStress) / float64(jumlah)
	persentaseSelesai := (float64(jumlahSelesai) / float64(jumlah)) * 100

	fmt.Printf("Rata-rata Stress		: %.2f\n", rataRataStress)
	fmt.Printf("Jumlah Tugas Selesai		: %.2f%%\n", persentaseSelesai)
	fmt.Println()
	rataRataStressMingguan := RataRataStressMingguan(arr,jumlah)
	fmt.Printf("Rata-rata Stress Mingguan	: %.2f\n", rataRataStressMingguan)
	fmt.Println()
	fmt.Print("Catatan Evaluasi: ")
	if persentaseSelesai == 100 {
		fmt.Println("Luar biasa! Tugas Selesai 100%")
	} else if rataRataStress >= 7.0 {
		fmt.Println("Tingkat stress tinggi, perlu penanganan lebih lanjut")
	} else if persentaseSelesai >= 50.0 {
		fmt.Println("Progress baik, 50% ","dari tugas sudah selesai")
	} else {
		fmt.Println("Perlu peningkatan, sebagian besar tugas belum selesai")
		
	}
	fmt.Println()

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
		fmt.Println("4. Hapus Pencapaian")
		fmt.Println("5. Urutkan Mood")
		fmt.Println("6. Urutkan Deadline")
		fmt.Println("7. Cari Berdasarkan...")
		fmt.Println("8. Tampilkan Statistik Pencapaian")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			TambahPencapaian(&data, &jumlahData)
		case 2:
			OutputPencapaian(&data, &jumlahData)
		case 3:
			UpdatePencapaian(&data, &jumlahData)
		case 4:
			HapusPencapaian(&data, &jumlahData)
		case 5:
			var subPilihan int
			fmt.Println("==Opsi Urutan Mood==")
			fmt.Println("1. Mood Tertinggi(Selection)")
			fmt.Println("2. Mood Terendah(Insertion)")
			fmt.Print("Pilih opsi: ")
			fmt.Scan(&subPilihan)
			fmt.Println()
			switch subPilihan {
			case 1:
				UrutkanMoodSelection(&data, jumlahData)
			case 2:
				UrutkanMoodInsertion(&data, jumlahData)
			default:
				fmt.Println("Opsi tidak valid")
				fmt.Println()
			}
		case 6:
			var subPilihan int
			fmt.Println("==Opsi Urutan Deadline==")
			fmt.Println("1. Deadline Terdekat(Insertion)")
			fmt.Println("2. Deadline Terjauh(Selection)")
			fmt.Print("Pilih opsi: ")
			fmt.Scan(&subPilihan)
			fmt.Println()
			switch subPilihan {
			case 1:
				UrutkanDeadlineTerdekat(&data, jumlahData)
			case 2:
				UrutkanDeadlineTerjauh(&data, jumlahData)
			default:
				fmt.Println("Opsi tidak valid")
				fmt.Println()
			}
		case 7:
			var subPilihan int
			fmt.Println("==Opsi Cari Berdasarkan==")
			fmt.Println("1. nama tugas(sequential)")
			fmt.Println("2. nama tugas(binary)")
			fmt.Println("3. tanggal deadline(sequential)")
			fmt.Println("4. tanggal deadline(binary)")
			fmt.Print("Pilih opsi: ")
			fmt.Scan(&subPilihan)
			fmt.Println()
			switch subPilihan {
			case 1:
				CariBerdasarkanTugas(&data, jumlahData)
			case 2:
				CariBerdasarkanTugasBinary(&data, jumlahData)
			case 3:
				CariBerdasarkanDeadline(&data, jumlahData)
			case 4:
				CariBerdasarkanDeadlineBinary(&data, jumlahData)
			default:
				fmt.Println("Opsi tidak valid")
				fmt.Println()
			}
		case 8:
			TampilkanStatistikPencapaian(&data, jumlahData)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
			fmt.Println()

		}
	}
}
