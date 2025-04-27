package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
		c int
		t []int
		a []int
	}
	type testCase struct {
		name string
		args args
		want []int
	}
	//再現性を維持したい場合は固定のシードを設定する
	rand.Seed(time.Now().UnixNano())
	var tests []testCase
	for i := 0; i < 1000; i++ {
		n := rand.Intn(2000)
		n++
		c := rand.Intn(1 << 30)
		var t, a []int
		for j := 0; j < n; j++ {
			ti := rand.Intn(3)
			ti++
			t = append(t, ti)
			a = append(a, rand.Intn(1<<30))
		}
		tests = append(tests, testCase{fmt.Sprintf("Case %03d", i+1), args{n, c, t, a}, solveHonestly(n, c, t, a)})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.c, tt.args.t, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
