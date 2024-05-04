package main

import (
	"reflect"
	"testing"
)

func Test_stack_carryOut(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"case 01", fields{[]int{0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				data: tt.fields.data,
			}
			s.carryOut()
			if !reflect.DeepEqual(s.data, []int{0}) {
				t.Errorf("stack.carryOut() = %v, want %v", s.data, []int{0})
			}
		})
	}
}

func Test_stack_liftManyBoxes(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{"case 01",
			fields{[]int{1, 2, 3, 4, 5}},
			args{3},
			[]int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				data: tt.fields.data,
			}
			if got := s.liftManyBoxes(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.liftManyBoxes() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(s.data, []int{1, 2, 3}) {
				t.Errorf("stack.liftManyBoxes() = %v, want %v", s.data, []int{1, 2, 3})
			}
		})
	}
}
