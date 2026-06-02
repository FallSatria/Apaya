package main 
import "fmt"

func main() {
	const maxN = 1000
	var a int
	fmt.Scan(&a)

	var abot [maxN]float64
	for i := 0; i < a; i++ {
		fmt.Scan(&abot[i])
	}
	if a == 0 {
		return
	}
	min := abot[0]
	max := abot[0]
	for i := 1; i < a; i++ {
		if abot[i] < min {
			min = abot[i]
		}
		if abot[i] > max {
			max = abot[i]
		}
	}
	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}