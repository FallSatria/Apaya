package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var target int
	fmt.Scan(&target)

	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == target {
			count++
		}
	}
	if count > 0 {
		fmt.Printf("%d ditemukan sebanyak %d kali\n", target, count)
	} else {
		fmt.Println("Tidak ditemukan")
	}
}