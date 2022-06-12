package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		k int
		a []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"sample1", args{2, 2, []int{2, 1}}, 3},
		{"sample2", args{3, 5, []int{1, 1, 1}}, 0},
		{"sample3", args{10, 998244353, []int{10, 9, 8, 7, 5, 6, 3, 4, 2, 1}}, 185297239},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := solve(tt.args.n, tt.args.k, tt.args.a); gotAns != tt.wantAns {
				t.Errorf("solve() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
