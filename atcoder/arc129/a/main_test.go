package main

import (
	"testing"
)

func Test_solveHonestly(t *testing.T) {
	type args struct {
		n int
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"in-01", args{2, 1, 2}, 1},
		{"in-02", args{10, 2, 19}, 10},
		{"hand-01", args{12, 30, 40}, 0},
		{"hand-02", args{20, 2, 19}, 8},
		{"hand-03", args{50, 30, 40}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveHonestly(tt.args.n, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("solveHonestly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestF(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"in-02-1", args{10, 19}, 10},
		{"in-02-2", args{10, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := F(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("F() = %v, want %v", got, tt.want)
			}
		})
	}
}
