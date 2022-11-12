package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		m int
		k int
		a []int
		b []int
		c []int
		e []int
	}
	type testCase struct {
		name string
		args args
		want int
	}
	var tests []testCase
	const max = 5
	for i := 0; i < 10000; i++ {
		n, m, k := rand.Intn(max-1)+2, rand.Intn(max)+1, rand.Intn(max)+1
		var a, b, c []int
		for i := 0; i < m; i++ {
			ai := rand.Intn(n)
			a = append(a, ai)
			for {
				bi := rand.Intn(n)
				if bi == ai {
					continue
				}
				b = append(b, bi)
				break
			}
			c = append(c, rand.Intn(100)+1)
		}
		var e []int
		for i := 0; i < k; i++ {
			e = append(e, rand.Intn(m))
		}
		//fmt.Println(n, m, k, a, b, c, e)
		tests = append(tests, testCase{
			fmt.Sprintf("Case %03d", i+1),
			args{n, m, k, a, b, c, e},
			solveCommentary(n, m, k, a, b, c, e),
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.m, tt.args.k, tt.args.a, tt.args.b, tt.args.c, tt.args.e); got != tt.want {
				t.Error(tt.args)
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
