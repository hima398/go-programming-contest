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
	"sort"
	"strconv"
	"strings"
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
	n int
	w int
	k int
	s int
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
		res = append(res, testCase{file.Name(), args{n, w, k, s, a, b, c, d}})
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

func Test_solve07(t *testing.T) {
	//h = h2
	var totalScores [][2]int
	for k := 1; k <= 10; k++ {
		fmt.Println("k = ", k)
		tests, js := readInput("./in")
		var totalScore int
		for i, tt := range tests {
			stdio = js[i]
			t.Run(tt.name, func(t *testing.T) {
				res, _, white, black, err := solve07(tt.args.n, tt.args.w, tt.args.k, tt.args.s, 100, k, 40, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
				if err != nil {
					t.Errorf(err.Error())
				}
				//writeOutput("./out/solve07", tt.name, res)
				//ExportImage("./img/solve07", strings.Replace(tt.name, ".txt", ".png", -1), tt.args.n, img)
				var score int
				for _, v := range res {
					score += v.p + tt.args.s
					totalScore += v.p + tt.args.s
				}
				scores[i][0] = i
				scores[i][2] = score
				scores[i][4] = white
				scores[i][5] = black
			})
		}
		totalScores = append(totalScores, [2]int{k, totalScore})
	}
	sort.Slice(totalScores, func(i, j int) bool {
		return totalScores[i][1] < totalScores[j][1]
	})
	for _, v := range totalScores {
		fmt.Println(v)
	}
	//fmt.Println("Solve07 = ", totalScore)
}

func Test_solve06(t *testing.T) {
	t.Skip()
	//h = h2
	tests, js := readInput("./in")
	var totalScore int
	for i, tt := range tests {
		stdio = js[i]
		t.Run(tt.name, func(t *testing.T) {
			res, img, err := solve06(tt.args.n, tt.args.w, tt.args.k, tt.args.s, 100, 2, 40, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if err != nil {
				t.Errorf(err.Error())
			}
			writeOutput("./out/solve06", tt.name, res)
			ExportImage("./img/solve06", strings.Replace(tt.name, ".txt", ".png", -1), tt.args.n, img)
			var score int
			for _, v := range res {
				score += v.p + tt.args.s
				totalScore += v.p + tt.args.s
			}
			scores[i][0] = i
			scores[i][1] = score
		})
	}
	fmt.Println("Solve06 = ", totalScore)
}

func Test_solve05(t *testing.T) {
	t.Skip()
	//h = h2
	tests, js := readInput("./in")
	var totalScore int
	for i, tt := range tests {
		stdio = js[i]
		t.Run(tt.name, func(t *testing.T) {
			res, _, err := solve05(tt.args.n, tt.args.w, tt.args.k, tt.args.s, 100, 2, 40, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			//res, _, err := solve04(tt.args.n, tt.args.w, tt.args.k, tt.args.s, 100, 500, 40, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if err != nil {
				t.Errorf(err.Error())
				//fmt.Println(err.Error())
			}
			//writeOutput("./out", tt.name, res)
			//ExportImage("./img", strings.Replace(tt.name, ".txt", ".png", -1), tt.args.n, img)
			for _, v := range res {
				totalScore += v.p + tt.args.s
			}
		})
	}

	fmt.Println("Solve05 = ", totalScore)
}

func Test_solve04(t *testing.T) {
	t.Skip()

	//h = h2
	var scores [][2]int
	//for p := 200; p <= 1000; p += 100 {
	//	fmt.Println("p = ", p)
	tests, js := readInput("./in")
	var totalScore int
	for i, tt := range tests {
		stdio = js[i]
		t.Run(tt.name, func(t *testing.T) {
			res, _, err := solve04(tt.args.n, tt.args.w, tt.args.k, tt.args.s, 100, 200, 40, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if err != nil {
				t.Errorf(err.Error())
				//fmt.Println(err.Error())
			}
			//writeOutput("./out", tt.name, res)
			//ExportImage("./img", strings.Replace(tt.name, ".txt", ".png", -1), tt.args.n, img)
			for _, v := range res {
				totalScore += v.p + tt.args.s
			}
		})
	}
	//scores = append(scores, [2]int{p, totalScore})
	//}
	//sort.Slice(scores, func(i, j int) bool {
	//	return scores[i][1] < scores[j][1]
	//})
	for _, v := range scores {
		fmt.Println(v[0], v[1])
	}
	//fmt.Println("Solve04 = ", totalScore)
}

func Test_solve035(t *testing.T) {
	t.Skip()
	//h = h2
	//var scores [][2]int
	//for p := 30; p <= 50; p++ {
	//fmt.Println("p = ", p)
	tests, js := readInput("./in")
	var totalScore int
	for i, tt := range tests {
		stdio = js[i]
		t.Run(tt.name, func(t *testing.T) {
			res, _, err := solve035(tt.args.n, tt.args.w, tt.args.k, tt.args.s, 40, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if err != nil {
				t.Errorf(err.Error())
				//fmt.Println(err.Error())
			}
			//solve04と比較するため、結果出力をコメントアウトしておく
			//writeOutput("./out", tt.name, res)
			//ExportImage("./img", strings.Replace(tt.name, ".txt", ".png", -1), tt.args.n, img)
			for _, v := range res {
				totalScore += v.p + tt.args.s
			}
		})
	}
	//scores = append(scores, [2]int{p, totalScore})
	//}
	//sort.Slice(scores, func(i, j int) bool {
	//	return scores[i][1] < scores[j][1]
	//})
	//for _, v := range scores {
	//	fmt.Println(v[0], v[1])
	//}
	fmt.Println("Solve035 = ", totalScore)
}

func Test_solve03(t *testing.T) {
	t.Skip()
	//h = h2
	tests, js := readInput("./in")
	var totalScore int
	for i, tt := range tests {
		stdio = js[i]
		t.Run(tt.name, func(t *testing.T) {
			res, img, err := solve03(tt.args.n, tt.args.w, tt.args.k, tt.args.s, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if err != nil {
				t.Errorf(err.Error())
				//fmt.Println(err.Error())
			}
			writeOutput("./out", tt.name, res)
			ExportImage("./img", strings.Replace(tt.name, ".txt", ".png", -1), tt.args.n, img)
			for _, v := range res {
				totalScore += v.p + tt.args.s
			}
		})
	}
	fmt.Println("Solve03 = ", totalScore)
}

func Test_solve02(t *testing.T) {
	t.Skip()
	//h = h2
	tests, js := readInput("./in")
	var totalScore int
	for i, tt := range tests {
		stdio = js[i]
		t.Run(tt.name, func(t *testing.T) {
			res, err := solve02(tt.args.n, tt.args.w, tt.args.k, tt.args.s, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if err != nil {
				t.Errorf(err.Error())
				//fmt.Println(err.Error())
			}
			writeOutput("./out", tt.name, res)
			for _, v := range res {
				totalScore += v.p + tt.args.s
			}
		})
	}
	fmt.Println("Solve02 = ", totalScore)
}
