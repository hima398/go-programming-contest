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
	d      int
	f1, r1 []string
	f2, r2 []string
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

		d := readInt()
		var f1, r1, f2, r2 []string
		for i := 0; i < d; i++ {
			f1 = append(f1, readString())
		}
		for i := 0; i < d; i++ {
			r1 = append(r1, readString())
		}
		for i := 0; i < d; i++ {
			f2 = append(f2, readString())
		}
		for i := 0; i < d; i++ {
			r2 = append(r2, readString())
		}
		res = append(res, testCase{file.Name(), args{d, f1, r1, f2, r2}})
	}
	return res
}

func writeOutput(dir, name string, n int, b0, b1 []int) {
	path := filepath.Join(dir, name)
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	fmt.Fprintln(fp, n)
	if len(b0) > 0 {
		fmt.Fprintf(fp, "%d", b0[0])
		for i := 1; i < len(b0); i++ {
			fmt.Fprintf(fp, " %d", b0[i])
		}
		fmt.Fprintln(fp)
	}
	if len(b1) > 0 {
		fmt.Fprintf(fp, "%d", b1[0])
		for i := 1; i < len(b1); i++ {
			fmt.Fprintf(fp, " %d", b1[i])
		}
		fmt.Fprintln(fp)
	}
}

func Test_solveHonestly(t *testing.T) {
	tests := readInput("./in")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, b0, b1 := solveHonestly(tt.args.d, tt.args.f1, tt.args.r1, tt.args.f2, tt.args.r2)
			writeOutput("./out/solveHonestly", tt.name, n, b0, b1)
		})
	}
}
