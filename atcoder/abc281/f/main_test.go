package main

import "testing"

func Test_dfs(t *testing.T) {
	type args struct {
		cur int
		a   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{1, []int{1, 1}}, 0},
		{"Case 2", args{1, []int{0, 1}}, 1},
		{"Case 3", args{2, []int{2, 3}}, 1},
		{"Case 4", args{3, []int{1, 3}}, 2},
		{"Case 5", args{2, []int{0, 1, 0, 0, 0, 1, 1, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.cur, tt.args.a); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
