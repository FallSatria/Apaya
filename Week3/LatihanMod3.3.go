package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalamLingkaran(cx, cy, r, x, y int) bool {
	dist := jarak(float64(x), float64(y), float64(cx), float64(cy))
	return dist <= float64(r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	
	fmt.Scan(&cx2, &cy2, &r2)
	
	fmt.Scan(&x, &y)

	masukLingkaran1 := didalamLingkaran(cx1, cy1, r1, x, y)
	masukLingkaran2 := didalamLingkaran(cx2, cy2, r2, x, y)

	if masukLingkaran1 && masukLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if masukLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if masukLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}