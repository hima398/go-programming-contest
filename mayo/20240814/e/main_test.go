package main

import (
	"testing"
)

func Test_check(t *testing.T) {
	type args struct {
		n int
		s []string
		r string
		c string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Input 1", args{5, []string{"AC..B", ".BA.C", "C.BA.", "BA.C.", "..CBA"}, "ABCBC", "ACAAB"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.n, tt.args.s, tt.args.r, tt.args.c); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
