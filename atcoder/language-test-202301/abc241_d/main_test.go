package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

type args struct {
	q int
	t []int
	x []int
	k []int
}

type testCase struct {
	name string
	args args
	//want []int
}

func writeOutput(dir, name string, args args) {
	path := filepath.Join(dir, name)
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	fmt.Fprintln(fp, args.q)
	for i := 0; i < args.q; i++ {
		if args.t[i] == 1 {
			fmt.Fprintln(fp, args.t[i], args.x[i])
		} else {
			fmt.Fprintln(fp, args.t[i], args.x[i], args.k[i])
		}
	}
}

func Test_solve(t *testing.T) {
	var tests []testCase
	tests = append(tests, testCase{"in-01", args{11, []int{1, 1, 1, 1, 3, 3, 3, 3, 2, 1, 2}, []int{20, 10, 30, 20, 15, 15, 15, 15, 100, 1, 100}, []int{0, 0, 0, 0, 1, 2, 3, 4, 5, 0, 5}}})

	{
		q := 2 * int(1e5)
		t, x, k := make([]int, q), make([]int, q), make([]int, q)
		for i := 0; i < q; i++ {
			t[i] = rand.Intn(3) + 1
			x[i] = rand.Intn(int(1e18)) + 1
			switch t[i] {
			case 2, 3:
				k[i] = rand.Intn(5) + 1
			}
		}
		tests = append(tests, testCase{"max-01", args{q, t, x, k}})
		writeOutput("./", "max-01.txt", args{q, t, x, k})
	}
	{
		q := 30000
		t, x, k := make([]int, q), make([]int, q), make([]int, q)
		for i := 0; i < q; i++ {
			t[i] = rand.Intn(3) + 1
			x[i] = rand.Intn(int(1e9)) + 1
			switch t[i] {
			case 2, 3:
				k[i] = rand.Intn(5) + 1
			}
		}
		tests = append(tests, testCase{"heavy-01", args{q, t, x, k}})
		writeOutput("./", "heavy-01.txt", args{q, t, x, k})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.q, tt.args.t, tt.args.x, tt.args.k)
			//if got := solve(tt.args.q, tt.args.t, tt.args.x, tt.args.k); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("solve() = %v, want %v", got, tt.want)
			//}
		})
	}
}
