package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		a int
		v int
		b int
		w int
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{1, 2, 3, 1, 3}, "YES"},
		{"case2", args{1, 2, 3, 2, 3}, "NO"},
		{"case3", args{1, 2, 3, 3, 3}, "NO"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.a, tt.args.v, tt.args.b, tt.args.w, tt.args.t); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
