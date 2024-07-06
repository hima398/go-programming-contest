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
		//want int
	}
	var tests []testCase
	numtest := 100
	for i := 0; i < numtest; i++ {
		n := 10
		var a []int
		for j := 0; j < n; j++ {
			a = append(a, rand.Intn(n))
		}
		tests = append(tests, testCase{fmt.Sprintf("Case %02d", i), args{n, a}})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := solve(tt.args.n, tt.args.a); got != tt.want {
			//	t.Errorf("solve() = %v, want %v", got, tt.want)
			//}
			fmt.Println(tt.name)
			fmt.Println(tt.args.n)
			fmt.Println(tt.args.a)
			solve(tt.args.n, tt.args.a)
		})
	}
}
