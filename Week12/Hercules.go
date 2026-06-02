package main

import "fmt"


func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		houses := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}
		selectionSort(houses)
		for j := 0; j < m; j++ {
			fmt.Print(houses[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
