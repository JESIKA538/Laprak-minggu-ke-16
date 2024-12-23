package main

import "fmt"

func main() {
	var n int
	var sum float64
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		suku := 1 / float64(2*i-1)
		if i%2 == 0 {
			suku = -suku
		}
		sum += suku
	}
	phi := 4 * sum
	fmt.Printf("N suku pertama: %d\n", n)
	fmt.Printf("Hasil PI: %.7f\n", phi)
}
