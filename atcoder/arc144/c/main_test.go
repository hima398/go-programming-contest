package main

import (
	"fmt"
	"math/rand"
	"reflect"
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
		want    []int
		wantErr bool
	}
	var tests []testCase

	tests = append(tests, testCase{"Case Hand 1", args{4, 1}, []int{2, 1, 4, 3}, false})
	tests = append(tests, testCase{"Case Hand 2", args{4, 2}, []int{3, 4, 1, 2}, false})
	tests = append(tests, testCase{"Case Hand 3", args{4, 3}, nil, true})
	tests = append(tests, testCase{"Case Hand 4", args{5, 1}, []int{2, 1, 4, 5, 3}, false})
	tests = append(tests, testCase{"Case Hand 5", args{5, 2}, []int{3, 4, 5, 1, 2}, false})
	tests = append(tests, testCase{"Case Hand 6", args{5, 3}, nil, true})
	tests = append(tests, testCase{"Case Hand 7", args{5, 4}, nil, true})
	tests = append(tests, testCase{"Case Hand 8", args{6, 1}, []int{2, 1, 4, 3, 6, 5}, false})
	tests = append(tests, testCase{"Case Hand 9", args{7, 1}, []int{2, 1, 4, 3, 6, 7, 5}, false})
	for i := 0; i < 20; i++ {
		n := rand.Intn(6)
		n += 2
		k := rand.Intn(n - 1)
		k++
		want, err := solveHonestly(n, k)
		wantErr := err != nil
		tests = append(tests, testCase{fmt.Sprintf("Case Random %d", i+1), args{n, k}, want, wantErr})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve(tt.args.n, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("n = %v, k = %v, solve() error = %v, wantErr %v", tt.args.n, tt.args.k, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("n = %v, k = %v, solve() = %v, want %v", tt.args.n, tt.args.k, got, tt.want)
			}
		})
	}
}
