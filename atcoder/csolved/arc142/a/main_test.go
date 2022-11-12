package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		k int
	}
	type testCase struct {
		name    string
		args    args
		wantAns int
	}
	var tests []testCase

	mx := int(1e5)
	for i := 0; i < 100; i++ {
		n, k := rand.Intn(mx+1), rand.Intn(mx+1)
		tests = append(tests, testCase{"Case " + strconv.Itoa(i+1) + ":", args{n, k}, solveHonestly(n, k)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := solve(tt.args.n, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("n = %v, k = %v, solve() = %v, want %v", tt.args.n, tt.args.k, gotAns, tt.wantAns)
			}
		})
	}
}
