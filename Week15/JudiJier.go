package main

import "fmt"
type Domino struct {
	sisi1 int  
	sisi2 int  
	nilai int  
	balak bool 
}
type Dominoes struct {
	kartu   [28]Domino 
	tersisa int        
}

func InisialisasiDomino(d *Dominoes) {
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			d.kartu[idx] = Domino{
				sisi1: i,
				sisi2: j,
				nilai: i + j,
				balak: i == j,
			}
			idx++
		}
	}
	d.tersisa = 28
}

func kocokKartu(d *Dominoes) {
	if d.tersisa <= 1 {
		return
	}
	
	seed := 12345
	
	for i := 0; i < d.tersisa; i++ {
		seed = (seed*1103515245 + 12345) % 2147483648
		j := seed % d.tersisa 
		if j < 0 {
			j = -j
		}
		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i]
	}
}
func ambilKartu(d *Dominoes) Domino {
	if d.tersisa == 0 {
		return Domino{} 
	}
	kartuYangDiambil := d.kartu[d.tersisa-1]
	d.tersisa--
	return kartuYangDiambil
}
func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.sisi1
	}
	return d.sisi2
}
func nilaiKartu(d Domino) int {
	return d.nilai
}
func galiKartu(d *Dominoes, patokan Domino) Domino {
	for d.tersisa > 0 {
		k := ambilKartu(d)
		if k.sisi1 == patokan.sisi1 || k.sisi1 == patokan.sisi2 ||
			k.sisi2 == patokan.sisi1 || k.sisi2 == patokan.sisi2 {
			return k
		}
	}
	return Domino{} 
}
func sepasangKartu(d1 Domino, d2 Domino) bool {
	return (d1.nilai + d2.nilai) == 12
}
var (
	pitaKarakter string
	indeks       int
	currentChar  byte
)
func start(input string) {
	pitaKarakter = input
	indeks = 0
	if len(pitaKarakter) > 0 {
		currentChar = pitaKarakter[indeks]
	}
}
func maju() {
	if !eop() {
		indeks++
		if indeks < len(pitaKarakter) {
			currentChar = pitaKarakter[indeks]
		}
	}
}
func eop() bool {
	if indeks >= len(pitaKarakter) {
		return true
	}
	return pitaKarakter[indeks] == '.'
}
func cc() byte {
	return currentChar
}
func prosesMesinKarakter(input string) {
	start(input)

	totalKarakter := 0
	jumlahA := 0
	jumlahLE := 0

	var prevChar byte = 0
	for !eop() {
		c := cc()
		totalKarakter++
		if c == 'A' {
			jumlahA++
		}
		if prevChar == 'L' && c == 'E' {
			jumlahLE++
		}

		prevChar = c
		maju()
	}
	frekuensiA := 0.0
	if totalKarakter > 0 {
		frekuensiA = float64(jumlahA) / float64(totalKarakter)
	}

	fmt.Println("--- Hasil Mesin Karakter ---")
	fmt.Printf("Teks Input          : %s\n", input)
	fmt.Printf("Total karakter      : %d\n", totalKarakter)
	fmt.Printf("Total huruf 'A'     : %d\n", jumlahA)
	// Cetak float menggunakan fmt
	fmt.Printf("Frekuensi huruf 'A' : %f\n", frekuensiA) 
	fmt.Printf("Total kata 'LE'     : %d\n", jumlahLE)
}

func main() {
	// --- Simulasi Domino ---
	fmt.Println("=== SIMULASI MESIN DOMINO ===")
	var setDomino Dominoes
	InisialisasiDomino(&setDomino)
	
	fmt.Printf("Jumlah kartu awal: %d\n", setDomino.tersisa)
	kocokKartu(&setDomino)
	
	kartu1 := ambilKartu(&setDomino)
	fmt.Printf("Ambil Kartu 1: (%d, %d) | Nilai: %d | Balak: %t\n", 
		gambarKartu(kartu1, 1), gambarKartu(kartu1, 2), nilaiKartu(kartu1), kartu1.balak)

	kartu2 := ambilKartu(&setDomino)
	fmt.Printf("Ambil Kartu 2: (%d, %d) | Nilai: %d | Balak: %t\n", 
		gambarKartu(kartu2, 1), gambarKartu(kartu2, 2), nilaiKartu(kartu2), kartu2.balak)

	isDuaBelas := sepasangKartu(kartu1, kartu2)
	fmt.Printf("Apakah total kedua kartu 12? %t\n", isDuaBelas)

	fmt.Printf("Mencari kartu (gali) yang cocok dengan (%d, %d)...\n", kartu1.sisi1, kartu1.sisi2)
	kartuCocok := galiKartu(&setDomino, kartu1)
	if kartuCocok.nilai != 0 || (kartuCocok.sisi1 == 0 && kartuCocok.sisi2 == 0) {
		fmt.Printf("Ditemukan kartu cocok: (%d, %d)\n", kartuCocok.sisi1, kartuCocok.sisi2)
	}
	fmt.Printf("Sisa kartu domino: %d\n\n", setDomino.tersisa)

	fmt.Println("=== SIMULASI MESIN KARAKTER ===")
	teks := "HALO ALAM LELE BERLEBIHAN."
	prosesMesinKarakter(teks)
}