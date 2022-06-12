package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		k int
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample1", args{4, 1, "abac"}, 2},
		{"sample2", args{10, 0, "aaaaaaaaaa"}, 1},
		{"sample3", args{6, 1, "abcaba"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.k, tt.args.s); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
