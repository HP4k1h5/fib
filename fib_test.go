package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	inputs := []int{0, 1, 2, 3, 4, 30, 50}
	expected := []int{0, 1, 1, 2, 3, 832040, 12586269025}
	for i, n := range inputs {
		actual := Fib(n)
		if actual != expected[i] {
			t.Errorf("expected %d to equal %d for input %d", actual, expected[i], n)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	inputs := []int{0, 1, 2, 3, 4, 30, 50}
	n := 0
	for i := 0; i < b.N; i++ {
		Fib(inputs[n%len(inputs)])
		n++
	}
}
