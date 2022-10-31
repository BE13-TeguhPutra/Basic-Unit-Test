package main

import "fmt"

func PrintLuas(alas, tinggi float64) float64 {
	// your code here

	var luas float64 = (alas * tinggi) / 2
	return luas
}

func main() {
	var alas float64 = 10
	var tinggi float64 = 10

	fmt.Println(PrintLuas(alas, tinggi))
}
