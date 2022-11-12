package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func generateString(n int) string {
	m := []string{"0", "1"}
	var res []string
	for i := 0; i < n; i++ {
		res = append(res, m[rand.Intn(2)])
	}
	return strings.Join(res, "")
}

func Test_solve(t *testing.T) {
	type args struct {
		n int
		s string
		t string
	}
	type testCase struct {
		name string
		args args
		want string
	}
	var tests []testCase
	for i := 0; i < 100; i++ {
		n := rand.Intn(10) + 1
		s, t := generateString(n), generateString(n)
		tests = append(tests, testCase{fmt.Sprintf("Case %02d", i+1), args{n, s, t}, solveHonestly(n, s, t)})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("s = %v, t = %v, solve() = %v, want %v", tt.args.s, tt.args.t, got, tt.want)
			}
		})
	}
}
