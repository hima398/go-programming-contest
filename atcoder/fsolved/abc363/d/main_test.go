package main

import (
	"fmt"
	"strconv"
	"testing"
)

var xs []string

func isPalindrome(s string) bool {
	n := len(s)
	ok := true
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		ok = ok && s[i] == s[j]
	}
	return ok
}

func solveHonestly(n int) {
	for x := 0; x <= n; x++ {
		s := strconv.Itoa(x)
		if isPalindrome(s) {
			xs = append(xs, s)
		}
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		n int
	}
	type testCase struct {
		name string
		args args
		want string
	}
	//tests := []testCase
	var tests []testCase
	tests = append(tests, testCase{"Sample 1", args{46}, "363"})
	tests = append(tests, testCase{"Sample 2", args{1}, "0"})
	tests = append(tests, testCase{"Sample 3", args{1000000000000000000}, "90000000000000000000000000000000009"})

	solveHonestly(int(1e8))

	for i := 1; i <= len(xs); i++ {
		tests = append(tests, testCase{fmt.Sprintf("Case %d", i), args{i}, xs[i-1]})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
