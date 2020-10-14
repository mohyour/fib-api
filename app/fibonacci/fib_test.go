package fibonacci

import (
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibonacci_400(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(100)
	}
}

func TestFibonacci(t *testing.T) {
	data := []struct {
		n    int
		want int64
	}{
		{-5, -5}, {0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := Fibonacci(d.n); got.Int64() != d.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}
