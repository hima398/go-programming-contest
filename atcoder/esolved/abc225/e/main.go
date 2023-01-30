package main

import (
	"bufio"
	"fmt"
	"math/cmplx"
	"os"
	"sort"
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
	ans := solve(n, x, y)
	PrintInt(ans)
}

func solve(n int, x, y []int) int {
	type fu struct {
		x1, y1 int
		x2, y2 int
		//s, t   float64
	}
	var fs []fu
	for i := range x {
		fs = append(fs, fu{x[i], y[i] - 1, x[i] - 1, y[i]})
	}
	sort.Slice(fs, func(i, j int) bool {
		return fs[i].y2*fs[j].x2 < fs[i].x2*fs[j].y2
	})

	ans := 0
	curX, curY := 0, 0
	for i := 0; i < n; i++ {
		if fs[i].x1*curY <= curX*fs[i].y1 {
			ans++
			curX, curY = fs[i].x2, fs[i].y2
		}
	}
	return ans
}

//角度順に区間スケジューリング(おそらく誤差で計算がおかしくなっている)
func solveByAngle(n int, x, y []int) int {
	type fu struct {
		s, t float64
	}
	var fs []fu
	for i := range x {
		z1 := complex(float64(x[i]), float64(y[i]-1))
		_, s := cmplx.Polar(z1)
		z2 := complex(float64(x[i]-1), float64(y[i]))
		_, t := cmplx.Polar(z2)
		fs = append(fs, fu{s, t})
	}
	sort.Slice(fs, func(i, j int) bool {
		return fs[i].t < fs[j].t
	})
	ans := 1
	cur := fs[0].t
	for i := 1; i < n; i++ {
		if fs[i].s >= cur {
			ans++
			cur = fs[i].t
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
