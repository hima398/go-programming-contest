package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		s string
	}
	type testCase struct {
		name    string
		args    args
		want    int
		wantErr bool
	}
	var tests []testCase

	tests = append(tests, testCase{"in01-1", args{3, "101"}, 1, false})
	tests = append(tests, testCase{"in01-2", args{6, "101101"}, 2, false})
	tests = append(tests, testCase{"in01-3", args{5, "11111"}, -1, true})
	tests = append(tests, testCase{"in01-4", args{6, "000000"}, 0, false})
	tests = append(tests, testCase{"in01-5", args{30, "111011100110101100101000000111"}, 8, false})
	//"11"を含む
	tests = append(tests, testCase{"hand01-1", args{2, "11"}, -1, true})
	tests = append(tests, testCase{"hand01-2", args{3, "011"}, -1, true})
	tests = append(tests, testCase{"hand01-3", args{3, "110"}, -1, true})
	tests = append(tests, testCase{"hand01-4", args{4, "1100"}, 2, false})
	tests = append(tests, testCase{"hand01-5", args{4, "0011"}, 2, false})
	tests = append(tests, testCase{"hand01-6", args{4, "0110"}, 3, false})
	tests = append(tests, testCase{"hand01-7", args{5, "11000"}, 2, false})
	tests = append(tests, testCase{"hand01-8", args{5, "00011"}, 2, false})
	tests = append(tests, testCase{"hand01-9", args{5, "01100"}, 2, false})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve(tt.args.n, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
