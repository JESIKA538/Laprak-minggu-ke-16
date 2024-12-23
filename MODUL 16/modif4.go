package main

import "fmt"

func main() {
	var n int
	var sum float64
	fmt.Scan(&n)

	result := 200002 
	var reresult float64

	for i := 1; i <= n; i++ {
		suku := 1 / float64(2*i-1)
		if i%2 == 0 {
			suku = -suku
		}
		sum += suku

		if i == result {
			reresult = 4 * sum
		}
	}

	phi := 4 * sum
	fmt.Printf("N suku pertama: %d\n", n)
	fmt.Printf("Hasil PI: %.10f\n", phi)
	fmt.Printf("Hasil PI: %.10f\n", reresult)
	fmt.Printf("Pada i ke: %d\n", result)
}
