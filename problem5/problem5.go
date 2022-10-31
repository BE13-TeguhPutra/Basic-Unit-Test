package main

import (
	"fmt"
	"math"
)

func PrintLuasPermukaan(radius, tinggi float64) float64 {
	// your code here

	var Luas_tabung float64 = 2 * math.Pi * radius * (radius + tinggi)
	return Luas_tabung
}

func main() {
	var T float64 = 20
	var r float64 = 4

	fmt.Println(PrintLuasPermukaan(r, T))
}
