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

const testCases = 100

var scores [][7]int

func TestMain(m *testing.M) {

	scores = make([][7]int, testCases)

	code := m.Run()

	/*
		for i := 0; i < 100; i++ {
			scores[i][3] = scores[i][1] - scores[i][2]
			scores[i][6] = scores[i][4] - scores[i][5]
		}
		sort.Slice(scores, func(i, j int) bool {
			return scores[i][3] < scores[j][3]
		})
		for _, v := range scores {
			fmt.Println(v)
		}
	*/
	os.Exit(code)
}

type args struct {
	n    int
	h, v []string
	d    [][]int
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
		readString := func() string {
			sc.Scan()
			return sc.Text()
		}
		n := readInt()
		var h []string
		for i := 0; i < n-1; i++ {
			h = append(h, readString())
		}
		var v []string
		for i := 0; i < n; i++ {
			v = append(v, readString())
		}
		d := make([][]int, n)
		for i := 0; i < n; i++ {
			d[i] = make([]int, n)
			for j := 0; j < n; j++ {
				d[i][j] = readInt()
			}
		}

		res = append(res, testCase{file.Name(), args{n, h, v, d}})
	}
	return res
}

func writeOutput(dir, name string, v any) {
	path := filepath.Join(dir, name)
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	fmt.Fprintf(fp, "%v\n", v)
}

func Test_solve02(t *testing.T) {
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := solve02(tt.args.n, tt.args.h, tt.args.v, tt.args.d)
			writeOutput("./out", tt.name, ans)
			score, _ := computeScore(tt.args.n, tt.args.d, ans)
			writeOutput("./score", tt.name, score)
		})
	}
}

func Test_solve00(t *testing.T) {
	t.Skip()
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := solve00(tt.args.n, tt.args.h, tt.args.v, tt.args.d)
			writeOutput("./out", tt.name, ans)
			score, _ := computeScore(tt.args.n, tt.args.d, ans)
			writeOutput("./score", tt.name, score)
		})
	}
}
