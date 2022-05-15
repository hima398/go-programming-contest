package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		n int
		a []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"sample1", args{5, []int{0, 7, 3, 0, 9}}, "Yes"},
		{"sample2", args{7, []int{4, 0, 9, 5, 11, 0, 0}}, "No"},
		{"same1", args{5, []int{3, 3, 3, 3, 3}}, "Yes"},
		{"same2", args{5, []int{3, 3, 3, 0, 0}}, "Yes"},
		{"allzero", args{5, []int{0, 0, 0, 0, 0}}, "Yes"},
		{"completed", args{3, []int{3, 1, 5}}, "Yes"},
		{"onlyone", args{10, []int{0, 0, 0, 0, 0, 0, 123, 0, 0, 0}}, "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.n, tt.args.a); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
