package main

import "testing"

func TestDiff(t *testing.T) {

	y := diff(2, 3)
	if y != 1 {
		t.Error("expected 1 got ", y)
	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diff(2, 3)
	}
}
