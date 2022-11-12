package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
	}
	type testCase struct {
		name string
		args args
		want int
	}
	var tests []testCase

	for i := 2; i < 100; i++ {
		tests = append(tests, testCase{fmt.Sprintf("N = %d", i), args{i}, solveHonestly(i)})
	}
	for i := 0; i < 100; i++ {
		n := rand.Intn(1003)
		n += 2
		tests = append(tests, testCase{fmt.Sprintf("Random:%d", i+1), args{n}, solveHonestly(n)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
