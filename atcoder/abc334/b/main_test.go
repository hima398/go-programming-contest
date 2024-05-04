package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		a int
		m int
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"例1", args{5, 3, -1, 6}, 3},
		{"例2", args{-2, 2, 1, 1}, 0},
		{"例3", args{-177018739841739480, 2436426, -80154573737296504, 585335723211047198}, 273142010859},

		{"負", args{0, 2, -5, -4}, 1},
		{"", args{0, 2, -4, -4}, 1},
		{"", args{0, 2, -4, -3}, 1},
		{"", args{0, 2, -3, -3}, 0},

		{"正", args{0, 2, 4, 5}, 1},
		{"", args{0, 2, 4, 4}, 1},
		{"", args{0, 2, 3, 4}, 1},
		{"", args{0, 2, 3, 3}, 0},

		{"両方", args{0, 2, -2, 2}, 3},
		{"", args{0, 2, -2, 1}, 2},
		{"", args{0, 2, -1, 2}, 2},
		{"", args{0, 2, -1, 1}, 1},
		{"", args{0, 2, -2, 0}, 2},
		{"", args{0, 2, -1, 0}, 1},
		{"", args{0, 2, 0, 0}, 1},
		{"", args{0, 2, 0, 1}, 1},
		{"", args{0, 2, 0, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.a, tt.args.m, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
