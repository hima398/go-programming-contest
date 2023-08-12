package main

import (
	"bufio"
	"fmt"
	"math"
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
	//ans := solveHonestly(n, x, y)
	ans := solve(n, x, y)
	PrintFloat64(ans)
}

func solve(n int, x, y []int) float64 {
	var ps []complex128
	for i := 0; i < n; i++ {
		ps = append(ps, complex(float64(x[i]), float64(y[i])))
	}
	sort.Slice(ps, func(i, j int) bool {
		_, t1 := cmplx.Polar(ps[i])
		_, t2 := cmplx.Polar(ps[j])
		return t1 < t2
	})

	var ans float64
	for i := 0; i < n; i++ {
		var cur complex128
		k := i
		for j := 0; j < n; j++ {
			cur += ps[k]
			ans = float64(math.Max(ans, cmplx.Abs(cur)))
			k = (k + 1) % n
		}
	}
	return ans
}

// O(N*2**N)なのでN=30くらいも計算困難
func solveHonestly(n int, x, y []int) float64 {
	var d2 int
	for pat := 0; pat < 1<<n; pat++ {
		var curX, curY int
		for i := 0; i < n; i++ {
			if (pat>>i)&1 > 0 {
				curX += x[i]
				curY += y[i]
			}
		}
		d2 = Max(d2, curX*curX+curY*curY)
	}
	ans := math.Sqrt(float64(d2))
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
