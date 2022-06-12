package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_solver(t *testing.T) {
	type args struct {
		n int
	}
	type testCase struct {
		name string
		args args
		want int
	}
	rand.Seed(time.Now().UnixNano())
	var tests []testCase
	for i := 0; i < 10; i++ {
		tests = append(tests, testCase{"", args{rand.Intn(int(1e6))}, 0})
	}
	tests = append(tests, testCase{"", args{11111}, 11111})
	tests = append(tests, testCase{"", args{111111}, 111111})
	tests = append(tests, testCase{"", args{200000}, 0})
	tests = append(tests, testCase{"", args{2000000}, 0})
	tests = append(tests, testCase{"", args{2000000000}, 0})
	tests = append(tests, testCase{"", args{489489489}, 0})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solver(tt.args.n); got != tt.want {
				t.Errorf("n = %v, solver() = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
