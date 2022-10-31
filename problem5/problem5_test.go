package main

import "testing"

func TestPrintLuasPermukaan(t *testing.T) {
	if PrintLuasPermukaan(4, 20) != 603.1857894892403 {
		t.Error("salah")
	}

}
