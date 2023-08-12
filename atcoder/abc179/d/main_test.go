package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		k int
		l []int
		r []int
	}
	type testCase struct {
		name string
		args args
		want int
	}
	var tests []testCase
	tests = append(tests, testCase{"Case 1", args{5, 2, []int{1, 3}, []int{1, 4}}, 4})
	tests = append(tests, testCase{"Case 2", args{5, 2, []int{3, 5}, []int{3, 5}}, 0})
	tests = append(tests, testCase{"Case 3", args{5, 1, []int{1}, []int{2}}, 5})
	tests = append(tests, testCase{"Case 4", args{60, 3, []int{5, 1, 10}, []int{8, 3, 15}}, 221823067})
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.k, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_solve(b *testing.B) {
	n := 2 * int(1e5)
	k := 1
	l := []int{1}
	r := []int{n}
	solve(n, k, l, r)
}
