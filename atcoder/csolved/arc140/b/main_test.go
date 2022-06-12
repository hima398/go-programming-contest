package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n int
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"sample1", args{6, "AARCCC"}, 2},
		{"sample2", args{5, "AAAAA"}, 0},
		{"sample3", args{9, "ARCARCARC"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
