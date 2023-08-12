package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	m := nextInt()
	t, p := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		t[i] = nextInt()
		if t[i] == 3 || t[i] == 4 {
			p[i] = nextInt()
		}
	}
	q := nextInt()
	var a, b []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(n, x, y, m, t, p, q, a, b)
	for _, v := range ans {
		PrintHorizonaly(v)
	}
}

type Vector [3]int

func NewVector(x, y int) *Vector {
	res := new(Vector)
	res[0] = x
	res[1] = y
	res[2] = 1
	return res
}

type Matrix [3][3]int

func NewMatrix() *Matrix {
	return new(Matrix)
}

// 時計回りに90度回転
//
//	0, 1, 0
//
// -1, 0, 0
//
//	0, 0, 1
func Rotate() *Matrix {
	res := NewMatrix()
	res[0][1] = 1
	res[1][0] = -1
	res[2][2] = 1
	return res
}

// 反時計回りに90度回転
//
// 0, -1, 0
//
// 1, 0, 0
//
// 0, 0, 1
func RotateReverse() *Matrix {
	res := NewMatrix()
	res[0][1] = -1
	res[1][0] = 1
	res[2][2] = 1
	return res
}

// x=pを軸を対称に反転する
// -1, 0, 2p
// 0, 1, 0
// 0, 0, 1
func ReflectX(p int) *Matrix {
	res := NewMatrix()
	res[0][0] = -1
	res[0][2] = 2 * p
	res[1][1] = 1
	res[2][2] = 1
	return res
}

// y=pを軸を対称に反転する
// 1, 0, 0
// 0, -1, 2p
// 0, 0, 1
func ReflectY(p int) *Matrix {
	res := NewMatrix()
	res[0][0] = 1
	res[1][1] = -1
	res[1][2] = 2 * p
	res[2][2] = 1
	return res
}

func (m *Matrix) mulVector(v *Vector) *Vector {
	res := NewVector(0, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res[i] += m[i][j] * v[j]
		}
	}
	return res
}

func (m *Matrix) mul(m2 *Matrix) *Matrix {
	res := NewMatrix()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				res[i][j] += m[i][k] * m2[k][j]
			}
		}
	}
	return res
}

func solve(n int, x, y []int, m int, t, p []int, q int, a, b []int) []Vector {
	var vs []*Vector
	for i := 0; i < n; i++ {
		vs = append(vs, NewVector(x[i], y[i]))
	}
	var ops []*Matrix
	for i := 0; i < m; i++ {
		switch t[i] {
		case 1:
			ops = append(ops, Rotate())
		case 2:
			ops = append(ops, RotateReverse())
		case 3:
			ops = append(ops, ReflectX(p[i]))
		case 4:
			ops = append(ops, ReflectY(p[i]))
		}
	}
	for i := 1; i < m; i++ {
		ops[i] = ops[i].mul(ops[i-1])
	}

	var ans []Vector
	for i := 0; i < q; i++ {
		if a[i] < 0 {
			ans = append(ans, *vs[b[i]])
		} else {
			ans = append(ans, *ops[a[i]].mulVector(vs[b[i]]))
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintHorizonaly(x Vector) {
	defer out.Flush()
	fmt.Fprintln(out, x[0], x[1])
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
