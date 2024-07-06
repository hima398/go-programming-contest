package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

type args struct {
	n int
	m int
	a []int
	b []int
	c []int
}
type testCase struct {
	name string
	args args
	want []int
	p    []int
}

func generateRandomTestCase() testCase {
	var res testCase
	n := 2 + rand.Intn(7)
	m := rand.Intn(n * (n - 1) / 2)

	var p []int
	for i := 0; i < n; i++ {
		p = append(p, i)
	}
	itr := rand.Intn(100)
	for cnt := 0; cnt < itr; cnt++ {
		i, j := rand.Intn(n), rand.Intn(n)
		p[i], p[j] = p[j], p[i]
	}

	var a, b, c []int
	for len(a) < m {
		x, y := rand.Intn(n), rand.Intn(n)
		if x == y {
			continue
		}
		if x < y {
			x, y = y, x
		}
		a = append(a, p[x])
		b = append(b, p[y])
		c = append(c, x-y)
	}
	res.args.n = n
	res.args.m = m
	res.args.a = a
	res.args.b = b
	res.args.c = c
	for i := range p {
		p[i]++
	}
	res.p = p
	return res
}

func Test_solve(t *testing.T) {
	var tests []testCase
	//tests = append(tests, testCase{})
	//tests = append(tests, testCase{})
	//tests = append(tests, testCase{})

	for cnt := 0; cnt < 100; cnt++ {
		tt := generateRandomTestCase()
		tt.name = fmt.Sprintf("Random %d", cnt+1)
		tt.want = solveHonestly(tt.args.n, tt.args.m, tt.args.a, tt.args.b, tt.args.c)
		tests = append(tests, tt)
	}
	for _, tt := range tests {
		fmt.Printf("tt: %v\n", tt)
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.m, tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testCase = %v, solve() = %v, want %v", tt, got, tt.want)
			}
		})
	}
}
