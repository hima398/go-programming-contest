package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(s, t string) int {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = Max(dp[i][j], dp[i-1][j], dp[i][j-1], dp[i-1][j-1]+1)
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := nextString(), nextString()
	ans := solve(s, t)
	PrintInt(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(p ...int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	sp := []int(p)
	res := sp[0]
	for i := 1; i < len(sp); i++ {
		res = max(res, sp[i])
	}
	return res
}
