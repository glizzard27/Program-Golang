//char-counter-version-1.0.go
//Next Update : char-counter-version-2.0.go

package main

import (
	"fmt"
)

func main() {
	// Meminta input dari pengguna
	var sentence string
	fmt.Println("-------------------------------------------")
	fmt.Print("Masukkan kalimat (Tanpa spasi): ")
	fmt.Scanf("%s", &sentence)
	fmt.Println("-------------------------------------------")

	// Menghitung banyaknya masing-masing karakter dalam kalimat
	charCount := make(map[rune]int)
	for _, char := range sentence {
		charCount[char]++
	}

	// Menampilkan hasilnya
	for char, count := range charCount {
		fmt.Printf("Jumlah karakter '%c': %d\n", char, count)
	}
	fmt.Println("-------------------------------------------")
}
