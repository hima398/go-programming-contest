package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

type args struct {
	n, m, d, k int
	u, v, w    []int
	x, y       []int
}
type testCase struct {
	name string
	args args
	want int
}

func readInput(dir string) []testCase {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var res []testCase
	for _, file := range files {
		//fmt.Println(filepath.Join(dir, f.Name()))
		filePath := filepath.Join(dir, file.Name())
		fp, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		sc := bufio.NewScanner(fp)
		sc.Split(bufio.ScanWords)

		readInt := func() int {
			sc.Scan()
			i, _ := strconv.Atoi(sc.Text())
			return int(i)
		}
		n, m, d, k := readInt(), readInt(), readInt(), readInt()
		var u, v, w []int
		for i := 0; i < m; i++ {
			u = append(u, readInt())
			v = append(v, readInt())
			w = append(w, readInt())
		}
		var x, y []int
		for i := 0; i < n; i++ {
			x = append(x, readInt())
			y = append(y, readInt())
		}
		res = append(res, testCase{file.Name(), args{n, m, d, k, u, v, w, x, y}, 1})
	}
	return res
}

func writeOutput(dir, name string, ans []int) {
	path := filepath.Join(dir, name)
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if len(ans) == 0 {
		return
	}

	fmt.Fprintf(fp, "%d", ans[0])
	for i := 1; i < len(ans); i++ {
		fmt.Fprintf(fp, " %d", ans[i])
	}
}

func Test_solve02(t *testing.T) {
	t.Skip()

	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, got := solve02(tt.args.n, tt.args.m, tt.args.d, tt.args.k, tt.args.u, tt.args.v, tt.args.w, tt.args.x, tt.args.y)
			writeOutput("./out", tt.name, ans)
			if got != tt.want {
				t.Errorf("solve02() = %v, want %v", got, tt.want)
			}
			fillPlan := true
			ds := make([]int, tt.args.d+1)
			for _, v := range ans {
				fillPlan = fillPlan && v > 0
				ds[v]++
			}
			if !fillPlan {
				t.Errorf("solve02() can not fill plan. ")
			}
			var isOver bool
			for i := 1; i <= tt.args.d; i++ {
				isOver = isOver || ds[i] > tt.args.k
			}
			if isOver {
				t.Errorf("solve02() is over.")
			}
		})
	}

}

func Test_solve03(t *testing.T) {
	t.Skip()

	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, got := solve03(tt.args.n, tt.args.m, tt.args.d, tt.args.k, tt.args.u, tt.args.v, tt.args.w, tt.args.x, tt.args.y)
			writeOutput("./out", tt.name, ans)
			if got != tt.want {
				t.Errorf("solve03() = %v, want %v", got, tt.want)
			}
			ok := true
			for _, v := range ans {
				ok = ok && v > 0
			}
			if !ok {
				t.Errorf("solve03() can not plan. ")
			}
		})
	}
}

func Test_solve04(t *testing.T) {
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, got := solve04(tt.args.n, tt.args.m, tt.args.d, tt.args.k, tt.args.u, tt.args.v, tt.args.w, tt.args.x, tt.args.y)
			writeOutput("./out", tt.name, ans)
			if got != tt.want {
				t.Errorf("solve04() = %v, want %v", got, tt.want)
			}
			fillPlan := true
			ds := make([]int, tt.args.d+1)
			for _, v := range ans {
				fillPlan = fillPlan && v > 0
				ds[v]++
			}
			if !fillPlan {
				t.Errorf("solve04() can not fill plan. ")
			}
			var isOver bool
			for i := 1; i <= tt.args.d; i++ {
				isOver = isOver || ds[i] > tt.args.k
			}
			if isOver {
				t.Errorf("solve04() is over.")
			}
		})
	}

}
