package main

import "fmt"

func PrintNama(nama string) string {
	// your code here
	var x string = fmt.Sprint("nama saya ", nama)

	return x
}

func main() {
	var nama string = "kobar"
	fmt.Println(PrintNama(nama))
}
