package main

import "testing"

func TestPrintLuas(t *testing.T) {
	// your code here
	if PrintLuas(10, 10) != 50 {
		t.Error("Perhitungan Luas Salah")
	}
}
