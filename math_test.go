package main

import "testing"

func TestSoma(t *testing.T) {
	result := soma(1, 2)
	if result != 3 {
		t.Errorf("Soma(1, 2) = %d; want 3", result)
	}
}
