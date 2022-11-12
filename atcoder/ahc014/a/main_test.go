package main

import "testing"

func Test_computeScore(t *testing.T) {
	type args struct {
		box [][]int
		d2  int
	}
	box := [][]int{
		//{1, 3, 3, 1, 2, 3, 1, 3, 3, 2},
		//{1, 1, 3, 3, 1, 3, 1, 3, 1, 1},
		//{2, 1, 2, 1, 2, 2, 2, 1, 1, 3},
		//{1, 3, 2, 1, 3, 2, 1, 2, 3, 2},
		//{3, 2, 3, 1, 2, 2, 2, 1, 2, 2},
		//{2, 1, 1, 3, 2, 1, 1, 2, 1, 2},
		//{1, 2, 2, 2, 3, 2, 2, 2, 3, 1},
		//{3, 2, 3, 1, 2, 3, 3, 3, 2, 2},
		//{1, 1, 3, 3, 1, 1, 1, 3, 2, 2},
		//{3, 1, 1, 1, 3, 1, 3, 2, 1, 2}}
		{3, 3, 1, 3, 2, 2, 1, 2, 2, 2},
		{3, 2, 3, 2, 1, 3, 3, 3, 1, 3},
		{2, 2, 1, 1, 1, 2, 2, 2, 1, 2},
		{2, 2, 2, 2, 1, 1, 2, 2, 3, 1},
		{1, 3, 1, 3, 1, 3, 1, 1, 2, 3},
		{1, 2, 3, 1, 1, 1, 3, 1, 1, 1},
		{3, 1, 1, 3, 1, 2, 3, 2, 2, 1},
		{2, 3, 2, 3, 3, 3, 2, 2, 3, 2},
		{1, 3, 2, 1, 3, 1, 2, 3, 1, 2},
		{3, 1, 1, 1, 3, 2, 1, 2, 2, 1},
	}
	type testCase struct {
		name string
		args args
		want int
	}
	var tests []testCase
	tests = append(tests, testCase{
		"Seed0",
		args{box, 3362},
		118977,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeScore(tt.args.box, tt.args.d2); got != tt.want {
				t.Errorf("computeScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
