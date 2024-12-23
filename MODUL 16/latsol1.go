package main

import "fmt"

func main() {
	var number, sum float64
	var count int
	for {
		fmt.Scan(&number)
		if number == 9999 {
			break
		}
		sum += number
		count++
	}
	if count > 0 {
		fmt.Printf("Rata-rata: %.2f\n", sum/float64(count))
	} else {
		fmt.Println("Tidak ada bilangan valid untuk dihitung rata-rata.")
	}
}

