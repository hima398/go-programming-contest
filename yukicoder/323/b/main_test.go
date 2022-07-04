package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		k int
	}
	type testCase struct {
		name string
		args args
		want int
	}
	var tests []testCase

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		n := rand.Intn(1000) + 1
		k := rand.Intn(n) + 1
		tc := testCase{fmt.Sprintf("Case %d:", i+1), args{n, k}, solveHonestly(n, k)}
		tests = append(tests, tc)
	}
	tests = append(tests, testCase{"Case 101:", args{1, 1}, 1})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
