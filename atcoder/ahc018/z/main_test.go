package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

type args struct {
	n int
	w int
	k int
	s int
	f [][]int
	a []int
	b []int
	c []int
	d []int
}

type testCase struct {
	name string
	args args
}

type testJudge struct {
	//y, x    int
	//isError bool
	n    int
	w, k int
	//水源の座標
	a, b []int
	//家の座標
	c, d []int
	r    [][]int
}

var isError bool
var gy, gx int

func (t testJudge) send(y, x, p int) {
	//掘削済みを指定していたらエラーを記憶しておく
	if t.r[y][x] <= 0 {
		isError = true
	}
	t.r[y][x] -= p
	//t.y = y
	//t.x = x
	gy, gx = y, x
}

func (t testJudge) receive() int {
	//エラーがあればそれを返す
	if isError {
		return -1
	}
	//すべてに水が通ったかチェック
	visited := make([][]bool, t.n)
	for i := range visited {
		visited[i] = make([]bool, t.n)
	}
	for i := 0; i < t.w; i++ {
		var q []Cell
		if t.r[t.a[i]][t.b[i]] <= 0 {
			q = append(q, Cell{t.a[i], t.b[i]})
			visited[t.a[i]][t.b[i]] = true
		}
		di := []int{-1, 0, 1, 0}
		dj := []int{0, -1, 0, 1}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if ni < 0 || ni >= t.n || nj < 0 || nj >= t.n {
					continue
				}
				if visited[ni][nj] {
					continue
				}
				if t.r[ni][nj] > 0 {
					continue
				}
				q = append(q, Cell{ni, nj})
				visited[ni][nj] = true
			}
		}
	}
	//fmt.Println(len(t.a), len(t.b), len(t.c), len(t.d))
	ok := true
	//for i := range t.a {
	//すべての家に水が通ったかの確認
	for i := 0; i < t.k; i++ {
		ok = ok && visited[t.c[i]][t.d[i]]
	}
	if ok {
		return 2
	}

	//壊れたかどうかのチェック
	//fmt.Println(t.y, t.x, t.r[gy][gx])
	if t.r[gy][gx] <= 0 {
		return 1
	} else {
		// r[t.y][t.x]>0
		return 0
	}
}

func readInput(dir string) ([]testCase, []testJudge) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var js []testJudge
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
		n, w, k, s := readInt(), readInt(), readInt(), readInt()
		r := make([][]int, n)
		for i := 0; i < n; i++ {
			r[i] = make([]int, n)
			for j := 0; j < n; j++ {
				r[i][j] = readInt()
			}
		}
		var a, b []int
		for i := 0; i < w; i++ {
			a = append(a, readInt())
			b = append(b, readInt())
		}
		var c, d []int
		for i := 0; i < k; i++ {
			c = append(c, readInt())
			d = append(d, readInt())
		}
		judge := testJudge{n: n, w: w, k: k, a: a, b: b, c: c, d: d, r: r}
		js = append(js, judge)
		res = append(res, testCase{file.Name(), args{n, w, k, s, r, a, b, c, d}})
	}
	return res, js
}

func writeOutput(dir, name string, ans []Excavation) {
	path := filepath.Join(dir, name)
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	for _, o := range ans {
		fmt.Fprintf(fp, "%d %d %d\n", o.y, o.x, o.p)
	}
}

func ExportImage(dir, name string, n int, field [][]int) {
	dest := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{n, n}})
	mx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mx = Max(mx, field[i][j])
		}
	}
	//fmt.Println("mx = ", mx)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			v := field[y][x] * 255 / mx
			v = Abs(v - 255)
			//fmt.Println(y, x, field[y][x], mx)
			dest.Set(x, y, color.Gray{uint8(v)})
		}
	}
	path := filepath.Join(dir, name)
	file, err := os.Create(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	png.Encode(file, dest)
}

func Test_solveWithField(t *testing.T) {
	//h = h2
	tests, js := readInput("./in")
	for i, tt := range tests {
		stdio = js[i]
		t.Run(tt.name, func(t *testing.T) {
			res, err := solveWithField(tt.args.n, tt.args.w, tt.args.k, tt.args.s, tt.args.f, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if err != nil {
				t.Errorf(err.Error())
			}
			writeOutput("./out", tt.name, res)
		})
	}
}
