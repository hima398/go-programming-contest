package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		x1 int
		x2 int
		x3 int
	}
	type testCase struct {
		name    string
		args    args
		want    int
		wantErr bool
	}
	var tests []testCase
	tests = append(tests, testCase{name: "In-01-1", args: args{2, 8, 8}})
	tests = append(tests, testCase{name: "In-01-2", args: args{1, 1, 1}})
	tests = append(tests, testCase{name: "In-01-3", args: args{5, 5, 10}})
	tests = append(tests, testCase{name: "In-01-4", args: args{10, 100, 1000}})
	for i := 0; i < 1000; i++ {
		x1, x2, x3 := rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)
		x1++
		x2++
		x3++
		tests = append(tests, testCase{name: fmt.Sprintf("Rand-%04d", i), args: args{x1, x2, x3}})
	}
	for i, tt := range tests {
		//fmt.Println("args = ", tt.args.x1, tt.args.x2, tt.args.x3)
		v, err := solveHonestly(tt.args.x1, tt.args.x2, tt.args.x3)
		if err != nil {
			tests[i].want = -1
			tests[i].wantErr = true
		} else {
			tests[i].want = v
			tests[i].wantErr = false
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve(tt.args.x1, tt.args.x2, tt.args.x3)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v, args = (%v, %v, %v)", err, tt.wantErr, tt.args.x1, tt.args.x2, tt.args.x3)
				return
			}
			if got != tt.want {
				t.Errorf("solve() = %v, want %v, args = (%v, %v, %v)", got, tt.want, tt.args.x1, tt.args.x2, tt.args.x3)
			}
		})
	}
}

func Test_solveHonestly(t *testing.T) {
	type args struct {
		x1 int
		x2 int
		x3 int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test",
			args:    args{238, 107, 496},
			want:    -1,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solveHonestly(tt.args.x1, tt.args.x2, tt.args.x3)
			if (err != nil) != tt.wantErr {
				t.Errorf("solveHonestly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solveHonestly() = %v, want %v", got, tt.want)
			}
		})
	}
}
