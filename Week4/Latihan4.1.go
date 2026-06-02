package main

import (
	"fmt"
)
func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}
func permutation(n int, r int, hasil *int) {
	var nFact, nrFact int
	factorial(n, &nFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / nrFact
}
func combination(n int, r int, hasil *int) {
	var nFact, rFact, nrFact int
	factorial(n, &nFact)
	factorial(r, &rFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / (rFact * nrFact)
}

func main() {
	var a, b, c, d int
	
	fmt.Scan(&a, &b, &c, &d)

	var permAC, kombAC int
	var permBD, kombBD int

	permutation(a, c, &permAC)
	combination(a, c, &kombAC)
	
	permutation(b, d, &permBD)
	combination(b, d, &kombBD)

	fmt.Println(permAC, kombAC)
	fmt.Println(permBD, kombBD)
}