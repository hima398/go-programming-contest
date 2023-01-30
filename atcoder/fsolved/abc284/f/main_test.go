package main

import (
	"testing"
)

func TestRollingHash_computeExcludedHash(t *testing.T) {
	type fields struct {
		p int
		n int
		s string
		w []int
		h []int
	}
	type args struct {
		l int
		r int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"",
			fields{
				2147483647,
				5,
				"abcde",
				[]int{1, int(1e2), int(1e4), int(1e6), int(1e8), int(1e10)},
				[]int{0, 1, 102, 10203, 1020304, 102030405},
			},
			args{0, 2},
			30405,
		},
		{
			"",
			fields{
				2147483647,
				5,
				"abcde",
				[]int{1, int(1e2), int(1e4), int(1e6), int(1e8), int(1e10)},
				[]int{0, 1, 102, 10203, 1020304, 102030405},
			},
			args{1, 3},
			10405,
		},
		{
			"",
			fields{
				2147483647,
				5,
				"abcde",
				[]int{1, int(1e2), int(1e4), int(1e6), int(1e8), int(1e10)},
				[]int{0, 1, 102, 10203, 1020304, 102030405},
			},
			args{2, 4},
			10205,
		},
		{
			"",
			fields{
				2147483647,
				5,
				"abcde",
				[]int{1, int(1e2), int(1e4), int(1e6), int(1e8), int(1e10)},
				[]int{0, 1, 102, 10203, 1020304, 102030405},
			},
			args{3, 5},
			10203,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := &RollingHash{
				p: tt.fields.p,
				n: tt.fields.n,
				s: tt.fields.s,
				w: tt.fields.w,
				h: tt.fields.h,
			}
			if got := hash.computeExcludedHash(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("RollingHash.computeExcludedHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
