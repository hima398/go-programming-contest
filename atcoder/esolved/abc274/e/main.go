package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	var p, q []int
	for i := 0; i < m; i++ {
		p = append(p, nextInt())
		q = append(q, nextInt())
	}
	ans := solve(n, m, x, y, p, q)
	PrintFloat64(ans)
}

func computeDist(x1, y1, x2, y2 int) float64 {
	xx := float64(x2) - float64(x1)
	yy := float64(y2) - float64(y1)
	return math.Sqrt(xx*xx + yy*yy)
}

func computeDist2(x1, y1, x2, y2 int) int {
	xx := x2 - x1
	yy := y2 - y1
	return xx*xx + yy*yy
}

func solve(n, m int, x, y, p, q []int) float64 {
	const INF = 1e18 + 1.0
	type pos struct {
		x, y int
	}
	var objects []pos
	for i := 0; i < m; i++ {
		objects = append(objects, pos{p[i], q[i]})
	}
	for i := 0; i < n; i++ {
		objects = append(objects, pos{x[i], y[i]})
	}
	mask := 1<<(n+m) - 1
	pat := make([][]int, 21)
	for k := 0; k <= mask; k++ {
		idx := bits.OnesCount(uint(k))
		pat[idx] = append(pat[idx], k)
	}
	for i := 0; i <= 20; i++ {
		sort.Ints(pat[i])
	}
	dp := make([][]float64, n+m)
	for i := 0; i < n+m; i++ {
		dp[i] = make([]float64, mask+1)
		for pattern := 0; pattern <= mask; pattern++ {
			dp[i][pattern] = INF
		}
	}
	for i := 0; i < n+m; i++ {
		dp[i][1<<i] = computeDist(0, 0, objects[i].x, objects[i].y)
	}
	for bitCount := 1; bitCount <= 20; bitCount++ {
		for _, cur := range pat[bitCount] {
			boost := math.Pow(2, float64(bits.OnesCount(uint((1<<m-1)&cur))))
			for from := 0; from < n+m; from++ {
				if cur>>from&1 == 0 {
					continue
				}
				for to := 0; to < n+m; to++ {
					if cur>>to&1 == 1 {
						continue
					}
					next := cur | 1<<to
					dp[to][next] = math.Min(dp[to][next], dp[from][cur]+computeDist(objects[from].x, objects[from].y, objects[to].x, objects[to].y)/boost)
				}
			}
		}
	}
	//fmt.Println(dp)
	ans := INF
	for i := 0; i < n+m; i++ {
		for pattern := 0; pattern <= mask; pattern++ {
			if bits.OnesCount(uint(pattern>>m)) != n {
				continue
			}
			boost := math.Pow(2, float64(bits.OnesCount(uint((1<<m-1)&pattern))))
			ans = math.Min(ans, dp[i][pattern]+computeDist(objects[i].x, objects[i].y, 0, 0)/boost)
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
