package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

// ポリオミノの個数
const m = 3

// グリッドの長さ
const n = 4

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	var p [][]string
	for i := 0; i < m; i++ {
		var pi []string
		for j := 0; j < n; j++ {
			pi = append(pi, nextString())
		}
		p = append(p, pi)
	}

	ok := solve(p)

	if ok {
		Print("Yes")
	} else {
		Print("No")
	}
}

type Polyomino struct {
	meta [][]int //元データ
	h, w int
}

func NewPolyomino(p []string) Polyomino {
	h, w := len(p), len(p[0])
	minI, maxI := h, 0
	minJ, maxJ := w, 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if p[i][j] == '#' {
				minI, maxI = Min(minI, i), Max(maxI, i)
				minJ, maxJ = Min(minJ, j), Max(maxJ, j)
			}
		}
	}
	res := Polyomino{h: maxI - minI + 1, w: maxJ - minJ + 1}
	meta := make([][]int, res.h)
	for i := range meta {
		meta[i] = make([]int, res.w)
		for j := 0; j < res.w; j++ {
			if p[minI+i][minJ+j] == '#' {
				meta[i][j] = 1
			} else {
				meta[i][j] = 0
			}
		}
	}
	res.meta = meta

	return res
}

func (p Polyomino) Rotate() Polyomino {
	res := Polyomino{h: p.w, w: p.h}
	meta := make([][]int, res.h)
	for i := range meta {
		meta[i] = make([]int, res.w)
	}
	for i := 0; i < p.h; i++ {
		for j := 0; j < p.w; j++ {
			ni, nj := j, res.w-1-i
			meta[ni][nj] = p.meta[i][j]
		}
	}
	res.meta = meta
	return res
}

func (p Polyomino) String() string {
	var res string
	for i := 0; i < p.h; i++ {
		res += fmt.Sprintln(p.meta[i])
	}
	return res
}

func solve(p [][]string) bool {
	var ps []Polyomino
	for i := 0; i < m; i++ {
		ps = append(ps, NewPolyomino(p[i]))
	}

	fillField := func(l int, si, sj []int, p []Polyomino) bool {
		var f [n][n]int
		for k := 0; k < l; k++ {
			for i := 0; i < p[k].h; i++ {
				for j := 0; j < p[k].w; j++ {
					if p[k].meta[i][j] > 0 {
						if f[si[k]+i][sj[k]+j] > 0 {
							return false
						} else {
							f[si[k]+i][sj[k]+j] = k + 1
						}
					}
				}
			}
		}
		ok := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ok = ok && f[i][j] > 0
			}
		}
		return ok
	}

	for k := 0; k < 4; k++ {
		ps[0] = ps[0].Rotate()
		for i := 0; i+ps[0].h <= n; i++ {
			for j := 0; j+ps[0].w <= n; j++ {
				for kk := 0; kk < 4; kk++ {
					ps[1] = ps[1].Rotate()
					for ii := 0; ii+ps[1].h <= n; ii++ {
						for jj := 0; jj+ps[1].w <= n; jj++ {
							for k3 := 0; k3 < 4; k3++ {
								ps[2] = ps[2].Rotate()
								for i3 := 0; i3+ps[2].h <= n; i3++ {
									for j3 := 0; j3+ps[2].w <= n; j3++ {
										if fillField(m, []int{i, ii, i3}, []int{j, jj, j3}, ps) {
											return true
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
