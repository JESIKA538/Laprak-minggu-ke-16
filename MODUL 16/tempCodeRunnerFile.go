package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var topping int
	fmt.Scan(&topping)

	result := 0
	for i := 0; i < topping; i++ {
		x, y := rand.Float64(), rand.Float64()
		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			result++
		}
	}
	fmt.Printf("Banyak Topping: %d\n", topping)
	fmt.Printf("Topping pada Pizza: %d\n", result)
	fmt.Printf("PI: %.10f\n", float64(result)*4/float64(topping))
}
