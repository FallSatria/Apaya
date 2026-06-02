package main

import "fmt"

func hitungPoin(level int) int {
	if level == 0 {
		return 0
	}
	var poin int
	fmt.Printf("Level ke-%d: ", 6-level)
	fmt.Scan(&poin)
	return poin + hitungPoin(level-1)
}

func main() {
	totalPoin := hitungPoin(5)

	bonus := 0
	if totalPoin >= 500 {
		bonus = 50
	}
	totalAkhir := totalPoin + bonus
	fmt.Printf("Total poin: %d\n", totalPoin)
	fmt.Printf("Bonus: %d\n", bonus)
	fmt.Printf("Total akhir: %d\n", totalAkhir)
}