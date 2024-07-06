package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		a int
		b int
		d []int
	}
	type testCase struct {
		name string
		args args
		want bool
	}
	var tests []testCase
	tests = append(tests, testCase{"Case 1", args{3, 2, 5, []int{1, 2, 9}}, true})
	tests = append(tests, testCase{"Case 2", args{2, 5, 10, []int{10, 15}}, false})
	tests = append(tests, testCase{"Case 3", args{4, 347, 347, []int{347, 700, 705, 710}}, true})

	const maxN = 20
	const maxV = 20
	for i := 0; i < 1000; i++ {
		name := fmt.Sprintf("Random %d", i+1)
		n := rand.Intn(maxN) + 1
		a, b := rand.Intn(maxV)+1, rand.Intn(maxV)+1
		var d []int
		for i := 0; i < n; i++ {
			d = append(d, rand.Intn(maxV)+1)
		}
		sort.Ints(d)
		tests = append(tests, testCase{name, args{n, a, b, d}, solveHonestly(n, a, b, d)})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.a, tt.args.b, tt.args.d); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
				t.Errorf("args = %v", tt.args)
			}
		})
	}
}
