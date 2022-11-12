package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"00", args{"KIHBR"}, "YES"},
		{"01", args{"KIHBRA"}, "YES"},
		{"02", args{"KIHBAR"}, "YES"},
		{"03", args{"KIHBARA"}, "YES"},
		{"04", args{"KIHABR"}, "YES"},
		{"05", args{"KIHABRA"}, "YES"},
		{"06", args{"KIHABAR"}, "YES"},
		{"07", args{"KIHABARA"}, "YES"},
		{"08", args{"AKIHBR"}, "YES"},
		{"09", args{"AKIHBRA"}, "YES"},
		{"10", args{"AKIHBAR"}, "YES"},
		{"11", args{"AKIHBARA"}, "YES"},
		{"12", args{"AKIHABR"}, "YES"},
		{"13", args{"AKIHABRA"}, "YES"},
		{"14", args{"AKIHABAR"}, "YES"},
		{"15", args{"AKIHABARA"}, "YES"},
		{"N00", args{"ABARA"}, "NO"},
		{"N01", args{"AAKIHABARA"}, "NO"},
		{"N02", args{"AKIHABAARA"}, "NO"},
		{"N03", args{"AKIHABARAA"}, "NO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
