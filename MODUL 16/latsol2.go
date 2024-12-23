package main

import "fmt"

func main() {
	var x string
	var n int

	fmt.Println("Masukkan string x:")
	fmt.Scan(&x)
	fmt.Println("Masukkan jumlah string n:")
	fmt.Scan(&n)

	var data []string

	fmt.Println("Masukkan", n, "string:")
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		data = append(data, str) 

	count := 0
	var positions []int

	for i, str := range data {
		if str == x {
			count++
			positions = append(positions, i+1) 
		}
	}

	// Menjawab pertanyaan
	fmt.Println("a. Apakah string x ada dalam kumpulan n data?")
	if count > 0 {
		fmt.Println("   Ya, string x ditemukan.")
	} else {
		fmt.Println("   Tidak, string x tidak ditemukan.")
	}

	fmt.Println("b. Pada posisi ke berapa string x ditemukan?")
	if count > 0 {
		fmt.Println("   Posisi:", positions)
	} else {
		fmt.Println("   String x tidak ditemukan.")
	}

	fmt.Println("c. Ada berapakah string x dalam kumpulan n data?")
	fmt.Printf("   String x muncul sebanyak: %d kali\n", count)

	fmt.Println("d. Adakah sedikitnya dua string x dalam n data?")
	if count >= 2 {
		fmt.Println("   Ya, ada sedikitnya dua string x.")
	} else {
		fmt.Println("   Tidak, tidak ada sedikitnya dua string x.")
	}

}
}
