package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

const testCases = 100

func TestMain(m *testing.M) {

	//scores = make([][7]int, testCases)

	code := m.Run()

	os.Exit(code)
}

type args struct {
	l int
	n int
	s int
	y []int
	x []int
	a []int
	f []int
}

type testCase struct {
	name string
	args args
}

type LocalSpace struct {
	n, l, s int
	y, x    []int
	a       []int
	f       []int
	p       [][]int
	//
	cnt int
}

var isError bool

func (local LocalSpace) place(p [][]int) {
	local.p = make([][]int, local.l)
	for i := range local.p {
		local.p[i] = make([]int, local.l)
		for j := range local.p[i] {
			local.p[i][j] = p[i][j]
		}
	}
}

func (local LocalSpace) measure(i, y, x int) int {
	res := Max(0, Min(1000, local.p[y][x]+local.f[local.cnt]))
	local.cnt++
	return res
}

func readInput(dir string) ([]testCase, []LocalSpace) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var localSpaces []LocalSpace
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
		l, n, s := readInt(), readInt(), readInt()
		y, x := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			y[i] = readInt()
			x[i] = readInt()
		}
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = readInt()
		}
		f := make([]int, 10005)
		for i := 0; i < 10000; i++ {
			f[i] = readInt()
		}
		space := LocalSpace{
			l: l, n: n, s: s, y: y, x: x, a: a, f: f}
		localSpaces = append(localSpaces, space)
		res = append(res, testCase{file.Name(), args{l: l, n: n, s: s, y: y, x: x}})
	}
	return res, localSpaces
}

/*
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
*/

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

func Test_solve01(t *testing.T) {
	tests, spaces := readInput("./in")
	for i, tt := range tests {
		anotherSpace = spaces[i]
		t.Run(tt.name, func(t *testing.T) {
			solve01(tt.args.l, tt.args.n, tt.args.s, tt.args.y, tt.args.x)
		})
	}
}
