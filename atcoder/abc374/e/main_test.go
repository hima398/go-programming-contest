package main

import (
	"testing"
)

func Test_check(t *testing.T) {
	type args struct {
		n int
		x int
		y int
		a []int
		p []int
		b []int
		q []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example 3", args{1, 1, 0, []int{1}, []int{10000000}, []int{1}, []int{10000000}}, true},
		{"example 3", args{1, 1, 1, []int{1}, []int{10000000}, []int{1}, []int{10000000}}, false},
		{"example 4", args{10, 7654321, 894742, []int{8, 5, 2, 7, 7, 4, 2, 1, 4, 6}, []int{6, 6, 4, 8, 9, 8, 2, 6, 2, 6}, []int{9, 4, 7, 9, 1, 9, 8, 2, 3, 5}, []int{1, 3, 9, 1, 6, 1, 9, 6, 4, 2}}, true},
		{"example 4", args{10, 7654321, 894743, []int{8, 5, 2, 7, 7, 4, 2, 1, 4, 6}, []int{6, 6, 4, 8, 9, 8, 2, 6, 2, 6}, []int{9, 4, 7, 9, 1, 9, 8, 2, 3, 5}, []int{1, 3, 9, 1, 6, 1, 9, 6, 4, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.n, tt.args.x, tt.args.y, tt.args.a, tt.args.p, tt.args.b, tt.args.q); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
