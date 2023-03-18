package main

import "testing"

func Benchmark_solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(2, 1000, []int{0, 0}, []int{0, 1}, []string{"B", "W"})
	}
}
