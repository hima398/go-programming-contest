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

	n, m := nextInt(), nextInt()
	ans := solve(n, m)
	PrintInt(ans)
}

func solve(n, m int) int {
	const p = 998244353
	dp := make([][2]int, n)
	dp[0][1] = m
	for i := 1; i < n; i++ {
		dp[i][0] += dp[i-1][0]*(m-2) + dp[i-1][1]*(m-1)
		dp[i][0] %= p

		dp[i][1] += dp[i-1][0]
		dp[i][1] %= p
	}
	return dp[n-1][0]
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
