package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		m []int
		p [][]int
		e [][]int
	}
	type testCase struct {
		name string
		args args
		want int
	}
	var tests []testCase
	ps := []int{2, 5, 7, 11}
	for ii := 0; ii < 100; ii++ {
		n := rand.Intn(11)
		m := make([]int, n)
		p, e := make([][]int, n), make([][]int, n)
		for i := 0; i < n; i++ {
			for jj := 0; jj < len(ps); jj++ {
				if rand.Intn(2)%2 == 0 {
					m[i]++
					p[i] = append(p[i], ps[jj])
					e[i] = append(e[i], rand.Intn(5))
				}
			}
			if m[i] == 0 {
				m[i]++
				p[i] = append(p[i], 2)
				e[i] = append(e[i], 1)
			}
		}
		tests = append(tests, testCase{fmt.Sprintf("Case:%03d", ii), args{n, m, p, e}, solveHonestly(n, m, p, e)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.m, tt.args.p, tt.args.e); got != tt.want {
				t.Error(tt.args)
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
