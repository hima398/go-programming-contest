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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, k := nextInt(), nextInt(), nextInt()
	ans := solve(n, m, k)
	PrintInt(ans)
}

func solve(n, m, k int) int {
	const p = 998244353

	//i番目まで見て数列の合計がjになる組み合わせの数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n*m+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for kk := 0; kk <= n*m; kk++ {
			for j := 1; j <= m; j++ {
				nextK := kk + j
				if nextK <= n*m {
					dp[i+1][nextK] += dp[i][kk]
					dp[i+1][nextK] %= p
				}
			}
		}
	}
	ans := 0
	//fmt.Println(dp[n])
	for kk := n; kk <= k; kk++ {
		ans += dp[n][kk]
		ans %= p
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
