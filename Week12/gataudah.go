package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int
	var num int
	for {
		_, err := fmt.Scan(&num)
		if err != nil || num < 0 {
			break 
		}
		arr = append(arr, num)
	}
	if len(arr) == 0 {
		return
	}
	insertionSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		if i < len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println() 
	isConstant := true
	var gap int

	if len(arr) > 1 {
		gap = arr[1] - arr[0]
		
		for i := 2; i < len(arr); i++ {
			if arr[i]-arr[i-1] != gap {
				isConstant = false
				break
			}
		}
	}
	if isConstant && len(arr) > 1 {
		fmt.Printf("Data berjarak %d\n", gap)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}