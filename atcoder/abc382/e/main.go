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

	n, x := nextInt(), nextInt()
	p := nextIntSlice(n)

	ans := solve(n, x, p)

	Print(ans)
}

func solve(n, x int, p []int) float64 {
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//pi目を引いてレアが出ない
			dp[i+1][j] += dp[i][j] * float64(100-p[i]) / 100.0
			dp[i+1][j+1] += dp[i][j] * float64(p[i]) / 100.0
		}
	}

	memo := make(map[int]float64)
	memo[0] = 0

	var f func(x int) float64
	f = func(x int) float64 {
		if v, found := memo[x]; found {
			return v
		}
		res := 1.0
		for j := 1; j <= n; j++ {
			y := Max(x-j, 0)
			res += f(y) * dp[n][j]
		}
		res /= 1 - dp[n][0]
		memo[x] = res
		return memo[x]
	}

	ans := f(x)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
