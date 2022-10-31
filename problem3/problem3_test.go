package main

import "testing"

func TestPrintNama(t *testing.T) {
	// your code here
	if PrintNama("teguh") != "nama saya teguh" {
		t.Error("Hasil cetak salah")
	}
}
