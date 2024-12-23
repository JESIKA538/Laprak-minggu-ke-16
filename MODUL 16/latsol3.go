package main

import "fmt"

func main() {
	var numDrops int
	fmt.Scan(&numDrops)
	resultA, resultB, resultC, resultD := 0, 0, 0, 0

	for i := 0; i < numDrops; i++ {
		if i%4 == 0 {
			resultA++ 
		} else if i%4 == 1 {
			resultB++ 
		} else if i%4 == 2 {
			resultC++
		} else {
			resultD++ 
		}
	}

	
	const Volume = 0.0001 
	curahA := float64(resultA) * Volume
	curahB := float64(resultB) * Volume
	curahC := float64(resultC) * Volume
	curahD := float64(resultD) * Volume

	// Menampilkan hasil
	fmt.Printf("Curah hujan untuk masing-masing daerah:\n")
	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", curahA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", curahB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", curahC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", curahD)
}
