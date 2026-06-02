package main

import "fmnt"

func fibonacci(n int) int {
	if  == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Print("n  |")
	for i := 0; i <= 10; i++ {
		fmt.Printf(" %-2d |", i) 
	}
	
	fmt.Println()

	fmt.Print("Sn |")
	for i := 0; i <= 10; i++ {
		fmt.Printf(" %-2d |", fibonacci(i))
	}
	
	fmt.Println() 
}