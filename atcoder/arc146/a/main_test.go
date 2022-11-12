package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		a []int
	}
	type testCase struct {
		name string
		args args
		want int
	}
	var tests []testCase
	tests = append(tests, testCase{"Sample 1", args{5, []int{1, 4, 3, 5, 8}}, 854})
	tests = append(tests, testCase{"Sample 2", args{8, []int{813, 921, 481, 282, 120, 900, 555, 409}}, 921900813})
	for i := 0; i < 1000; i++ {
		//n := rand.Intn(97) + 3
		n := 4
		var a []int
		for j := 0; j < n; j++ {
			a = append(a, rand.Intn(int(1e6)-1)+1)
		}
		tests = append(tests, testCase{fmt.Sprintf("Random %4d", i+1), args{n, a}, solveHonestly(n, a)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.a); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
