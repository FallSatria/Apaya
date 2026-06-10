package main

import "fmt"

var seed uint32 = 123456789

func angkaAcak() int {
	seed = (1103515245*seed + 12345) % 2147483648
	return int(seed)
}

type domino struct {
	sisi1 int
	sisi2 int
	Nilai int
	Balak bool
}

type Dominoes struct {
	kartu [28]domino
	Sisa int
}

func InitDominoes() *Dominoes {
	d := &Dominoes{}
	index := 0
	for i := 0: i <=6 ; i++ {
		for j := i; j <= 6; j++ {
			d.kartu[index] = Domino{
				sisi1: i,
				sisi2: j,
				Nilai: i + j,
				Balak: i == j,
			}
			index++
		
			
		}
	}
	d.Sisa = 28
	return d
}

func kocokKartu(d *Dominoes) {
	for i := d.Sisa - 1; i > 0; i-- {
		j := angkaAcak() % (i + 1)
		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i]
	}
}

func ambilKartu(d *Dominoes) domino {
	if d.Sisa <= 0 {
		return Domino{} 
	}
	kartu := d.kartu[d.Sisa-1]
	d.Sisa--
	return kartu
}

func nilaiKartu(d domino) int {
	return d.Nilai
}

func galiKartu(d domino) int {
	retrun d.Nilai
}

func galiKartu(d *dominoes, target Domino) int {
	for d.Sisa > 0 {
		kartu := ambilKartu(d)
		if kartu.sisi1 == target.sisi1 || kartu.sisi1 == target.sisi2 {
			kartu.sisi2 == target.sisi1 || kartu.sisi2 == target.sisi2 {
				return kartu
		}
	}
	return Domino{}
}

func sepasangKartu(d1 domino, d2 domino) bool {
	return (nilaiKartu(d1) + nilaiKartu(d2)) == 12
}

type MesinKarakter struct {
	Pita string
	Posisi int
	Karakter string
}

func (m *MesinKarakter) start(input string) {
	m.Pita = input
	m.Posisi = 0
	if len(m.Pita) > 0 {
		m.Karakter = string(m.Pita[m.Posisi])
	}
} 

func (m *MesinKarakter) next() {
	m.Posisi++
	if m.Posisi < len(m.Pita) {
		m.Karakter = string(m.Pita[m.Posisi])
	} 
}

func (m *MesinKarakter) isEnd() bool {
	if m.Posisi >= len(m.Pita) {
		return true
	}
	return string(m.Pita[m.Posisi]) == "."
}

func jalankanAnalisis(input string) {
	mesin := &MesinKarakter{}
	mesin.start(input)

	totalKarakter := 0
	totalA := 0
	totalLE := 0

	karakterSebelumnya := ""

	for !mesin.isEnd() {
		cc := mesin.cc()
		totalKarakter++

		if cc == "A" {
			totalA++
		}

		if karakterSebelumnya == "L" && cc == "E" {
			totalLE++
		}

		karakterSebelumnya = cc
		mesin.next()
	}

	frekuensiA := 0
	if totalKarakter > 0 {
		frekuensiA = (totalA * 100) / totalKarakter
	}
	fmt.Printf("Pita Dibaca: %s\n", input)
	fmt.Printf("Total Karakter: %d\n", totalKarakter)
	fmt.Printf("Jumlah A: %d\n", totalA)
	fmt.Printf("Frekuensi A: %d%%\n", frekuensiA)
	fmt.Printf("Jumlah LE: %d\n", totalLE)

}

