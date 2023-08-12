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
	ans := solve(n)
	PrintInt(ans)
}

func solve(n int) int {
	const p = 998244353
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 10)
	}
	for j := 1; j < 10; j++ {
		dp[1][j] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < 10; j++ {
			dp[i+1][j] += dp[i][j]
			dp[i+1][j] %= p
			pj, nj := j-1, j+1
			if pj >= 0 {
				dp[i+1][pj] += dp[i][j]
				dp[i+1][pj] %= p
			}
			if nj < 10 {
				dp[i+1][nj] += dp[i][j]
				dp[i+1][nj] %= p
			}
		}
	}
	var ans int
	for j := 1; j < 10; j++ {
		ans += dp[n][j]
		ans %= p
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
