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
	n int
	m int
	k int
	x []int
	y []int
	u []int
	v []int
	w []int
	a []int
	b []int
}
type testCase struct {
	name string
	args args
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
		n, m, k := readInt(), readInt(), readInt()
		var x, y []int
		for i := 0; i < n; i++ {
			x = append(x, readInt())
			y = append(y, readInt())
		}
		var u, v, w []int
		for i := 0; i < m; i++ {
			u = append(u, readInt()-1)
			v = append(v, readInt()-1)
			w = append(w, readInt())
		}
		var a, b []int
		for i := 0; i < k; i++ {
			a = append(a, readInt())
			b = append(b, readInt())
		}
		res = append(res, testCase{file.Name(), args{n, m, k, x, y, u, v, w, a, b}})
	}
	return res
}

func writeOutput(dir, name string, p, c []int) {
	path := filepath.Join(dir, name)
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if len(p) > 0 {
		fmt.Fprintf(fp, "%d", p[0])
		for i := 1; i < len(p); i++ {
			fmt.Fprintf(fp, " %d", p[i])
		}
	}
	fmt.Fprintln(fp)

	if len(c) > 0 {
		fmt.Fprintf(fp, "%d", c[0])
		for i := 1; i < len(c); i++ {
			fmt.Fprintf(fp, " %d", c[i])
		}
	}
	fmt.Fprintln(fp)
}

func Test_solve05(t *testing.T) {
	var totalScore int
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotC := solve05(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.u, tt.args.v, tt.args.w, tt.args.a, tt.args.b)
			writeOutput("./out/", tt.name, gotP, gotC)
			score, ok := computeScore(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.w, tt.args.a, tt.args.b, gotP, gotC)
			totalScore += score
			if !ok {
				fmt.Println(tt.name + " is not covered.")
			}
		})
	}
	fmt.Println("Case 5:totalScore = ", totalScore)
}

func Test_solve04(t *testing.T) {
	t.Skip()
	var totalScore int
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotC := solve04(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.u, tt.args.v, tt.args.w, tt.args.a, tt.args.b)
			writeOutput("./out/", tt.name, gotP, gotC)
			score, ok := computeScore(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.w, tt.args.a, tt.args.b, gotP, gotC)
			totalScore += score
			if !ok {
				fmt.Println(tt.name + " is not covered.")
			}
		})
	}
	fmt.Println("Case 4:totalScore = ", totalScore)
}

func Test_solve03(t *testing.T) {
	t.Skip()
	var totalScore int
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotC := solve03(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.u, tt.args.v, tt.args.w, tt.args.a, tt.args.b)
			writeOutput("./out/", tt.name, gotP, gotC)
			score, ok := computeScore(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.w, tt.args.a, tt.args.b, gotP, gotC)
			totalScore += score
			if !ok {
				fmt.Println(tt.name + " is not covered.")
			}
		})
	}
	fmt.Println("Case 3:totalScore = ", totalScore)
}

func Test_solve02(t *testing.T) {
	t.Skip()
	var totalScore int
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotC := solve02(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.u, tt.args.v, tt.args.w, tt.args.a, tt.args.b)
			writeOutput("./out/", tt.name, gotP, gotC)
			score, ok := computeScore(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.w, tt.args.a, tt.args.b, gotP, gotC)
			totalScore += score
			if !ok {
				fmt.Println(tt.name + " is not covered.")
			}
		})
	}
	fmt.Println("Case 2:totalScore = ", totalScore)
}

func Test_solve01(t *testing.T) {
	t.Skip()
	var totalScore int
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotC := solve01(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.u, tt.args.v, tt.args.w, tt.args.a, tt.args.b)
			writeOutput("./out/", tt.name, gotP, gotC)
			score, ok := computeScore(tt.args.n, tt.args.m, tt.args.k, tt.args.x, tt.args.y, tt.args.w, tt.args.a, tt.args.b, gotP, gotC)
			totalScore += score
			if !ok {
				fmt.Println(tt.name + " is not covered.")
			}
		})
	}
	fmt.Println("Case 1:totalScore = ", totalScore)
}
