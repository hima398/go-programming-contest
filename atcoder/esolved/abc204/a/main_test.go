package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{0, 0}, 0},
		{"", args{1, 1}, 1},
		{"", args{2, 2}, 2},
		{"", args{0, 1}, 2},
		{"", args{1, 2}, 0},
		{"", args{0, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
