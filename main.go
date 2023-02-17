package main

import "fmt"

func main() {
	fmt.Println("Tes")

	// Pembuatan hitung
	a := 2
	b := 4
	c := a + b

	fmt.Println("Hasil perhitungan = ", c)

	// perkalian
    g := 7
    h := 3
    i := g * h
    fmt.Println("Hasil perkalian = ", i)

	// Luas Lingkaran
	var r float64
	var pi float64 = 3.14

    fmt.Print("Masukkan jari-jari lingkaran: ")
    fmt.Scanln(&r)

    luas := pi * r * r

    fmt.Printf("Luas lingkaran adalah ", luas)
}