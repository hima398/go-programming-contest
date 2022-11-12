package main

import "testing"

func Test_firstSolve(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstSolve(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("firstSolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
