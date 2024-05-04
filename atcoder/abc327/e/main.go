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

	n := nextInt()
	p := nextIntSlice(n)

	ans := firstsolve(n, p)
	//ans := solve(n, p)
	Print(ans)
}

func solve(n int, p []int) float64 {
	const INF = float64(1 << 60)
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
		for j := range dp[i] {
			dp[i][j] = -INF
		}
	}
	//fmt.Println(w)
	for i := 1; i <= n; i++ {
		//i個目を選ばない場合
		for k := 1; k <= i; k++ {
			dp[i][k] = math.Max(dp[i][k], dp[i-1][k])
		}
		//i個目を選ぶ場合
		for k := 1; k <= i; k++ {
			var next float64
			if dp[i-1][k-1] > -INF {
				next = dp[i-1][k-1]
			}
			next *= 0.9
			next += float64(p[i-1])
			dp[i][k] = math.Max(dp[i][k], next)
		}
	}
	w := make([]float64, n+1)
	w[1] = 1.0
	for k := 2; k <= n; k++ {
		w[k] = 0.9*w[k-1] + 1.0
	}

	ans := -INF

	for k := 1; k <= n; k++ {
		ans = math.Max(ans, dp[n][k]/w[k]-1200/math.Sqrt(float64(k)))
	}
	return ans
}

func firstsolve(n int, p []int) float64 {
	const INF = float64(1 << 60)
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
		for j := range dp[i] {
			dp[i][j] = -INF
		}
	}
	w := make([]float64, n+1)
	w[1] = 1.0
	for k := 2; k <= n; k++ {
		w[k] = 0.9*w[k-1] + 1.0
	}
	for i := 1; i <= n; i++ {
		//i個目を選ばない場合
		for k := 0; k <= i; k++ {
			dp[i][k] = dp[i-1][k]
		}
		//i個目を選ぶ場合
		for k := 1; k <= i; k++ {
			var next float64
			if dp[i-1][k-1] > -INF {
				next = dp[i-1][k-1] + 1200.0/math.Sqrt(float64(k-1))
			}
			next *= w[k-1]
			next *= 0.9
			next += float64(p[i-1])
			next /= w[k]
			next -= 1200.0 / math.Sqrt(float64(k))
			dp[i][k] = math.Max(dp[i][k], next)
		}
	}
	ans := -INF
	for k := 1; k <= n; k++ {
		ans = math.Max(ans, dp[n][k])
	}
	return ans
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
