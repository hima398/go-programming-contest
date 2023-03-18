package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		a int
		b int
		c int
		d int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{4, 0, 2, 1, 0}, true},
		{"", args{4, 3, 0, 0, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
