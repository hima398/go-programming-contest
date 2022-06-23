package main

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
	}
	judge := func(ans [][]int) bool {
		n := len(ans)
		for i := 1; i < n-1; i++ {
			for j := 1; j < n-1; j++ {
				var a, b int
				part := make([][]int, 3)
				for ii := 0; ii < 3; ii++ {
					part[ii] = make([]int, 3)
				}
				for ii := i - 1; ii <= i+1; ii++ {
					for jj := j - 1; jj <= j+1; jj++ {
						if ii == i && jj == j {
							continue
						}
						if ans[ii][jj] > ans[i][j] {
							a++
						} else {
							b++
						}
						part[i-ii+1][j-jj+1] = ans[ii][jj]
					}
				}
				if a == b {
					PrintVertically(part)
					return false
				}
			}
		}
		return true
	}

	type testCase struct {
		name string
		args args
	}
	var tests []testCase
	for i := 2; i <= 500; i++ {
		tests = append(tests, testCase{fmt.Sprintf("Case %d:", i), args{i}})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.args.n)
			ok := judge(got)
			if !ok {
				t.Errorf("n = %v, This case is not satisfied conditions.", tt.args.n)
			}
		})
	}
}
