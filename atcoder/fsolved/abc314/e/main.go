package main

import (
	"bufio"
	"fmt"
	"math"
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

	n, m := nextInt(), nextInt()
	var c, p []int
	var s [][]int
	for i := 0; i < n; i++ {
		c = append(c, nextInt())
		p = append(p, nextInt())
		s = append(s, nextIntSlice(p[i]))
	}

	ans := solve(n, m, c, p, s)

	Print(ans)
}

func solve(n, m int, c, p []int, s [][]int) float64 {
	//kポイント持っているとき、mポイント以上獲得するコストの期待値
	dp := make([]float64, m+1)
	for k := m - 1; k >= 0; k-- {
		dp[k] = math.MaxFloat64
		for i := range s {
			var sum float64
			var nz float64
			for j := range s[i] {
				if s[i][j] > 0 {
					sum += dp[Min(k+s[i][j], m)]
				} else {
					nz++
				}
			}
			sum = (sum + float64(p[i]*c[i])) / (float64(p[i]) - nz)
			dp[k] = math.Min(dp[k], sum)
		}
	}

	return dp[0]
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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
