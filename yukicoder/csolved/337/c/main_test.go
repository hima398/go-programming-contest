package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n int
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample 1", args{4, 1, 4}, 2},
		{"sample 2", args{6, 3, 5}, 1},
		{"sample 3", args{4, 3, 2}, 0},
		{"sample 4", args{31415, 9265, 3589}, 889048175},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
