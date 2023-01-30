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
	var a, b []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(n, a, b)
	PrintVertically(ans)
}

func solve(n int, a, b []int) []int {
	const p = 998244353

	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	var dfs func(cur, par int) [][]int
	dfs = func(cur, par int) [][]int {
		dp := make([][]int, 2)
		for i := range dp {
			dp[i] = make([]int, 2)
		}
		dp[0][0] = 1
		dp[1][1] = 1
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			res := dfs(next, cur)
			ndp := make([][]int, 2)
			for i := range ndp {
				ndp[i] = make([]int, len(dp[0])+len(res[0])-1)
			}
			for i := range dp[0] {
				for j := range res[0] {
					ndp[0][i+j] += dp[0][i] * (res[0][j] + res[1][j])
					ndp[0][i+j] %= p

					ndp[1][i+j] += dp[1][i] * res[0][j]
					ndp[1][i+j] %= p

					if i+j > 0 {
						ndp[1][i+j-1] += dp[1][i] * res[1][j]
						ndp[1][i+j-1] %= p
					}
				}
			}
			dp, ndp = ndp, dp
		}
		return dp
	}
	dp := dfs(0, -1)
	var ans []int
	for i := 1; i <= n; i++ {
		v := (dp[0][i] + dp[1][i]) % p
		ans = append(ans, v)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
