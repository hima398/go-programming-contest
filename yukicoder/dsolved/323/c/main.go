package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, m, lt int, s, t []int) int {
	const p = 998244353
	//i日目(0<=i<=T)の都市j(0<=j<n)の人数
	dp := make([][]int, lt+1)
	for i := 0; i <= lt; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 0
	}
	dp[0][0] = 1

	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[s[i]] = append(e[s[i]], t[i])
		e[t[i]] = append(e[t[i]], s[i])
	}
	var dfs func(i, idx, par int) int
	dfs = func(i, idx, par int) int {
		if dp[i][idx] >= 0 {
			return dp[i][idx]
		}
		var s int
		if par >= 0 {
			s += dp[i-1][par]
		}
		for _, next := range e[idx] {
			if next == par {
				continue
			}
			s += dfs(i-1, next, idx)
		}
		s %= p
		dp[i][idx] = s
		return dp[i][idx]
	}

	for i := 1; i <= lt; i++ {
		dfs(i, 0, -1)
	}
	//fmt.Println(dp)
	return dp[lt][0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, lt := nextInt(), nextInt(), nextInt()
	var s, t []int
	for i := 0; i < m; i++ {
		s = append(s, nextInt())
		t = append(t, nextInt())
	}

	ans := solve(n, m, lt, s, t)

	PrintInt(ans)
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
