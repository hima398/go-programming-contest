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

	n, li := nextInt(), nextInt()
	var s, a []int
	for i := 0; i < n; i++ {
		s = append(s, nextInt())
		a = append(a, nextInt())
	}
	ans := solve(n, li, s, a)
	PrintInt(ans)
}

func solve(n, li int, s, a []int) int {
	//i杯目のラーメンまで見て、トータルの塩の濃さがjで食べられる味の濃さの最大値
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, li+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= li; j++ {
			//i杯目のラーメンを食べない
			dp[i][j] = dp[i-1][j]
			//i杯目のラーメンを食べる
			nextS := j + s[i-1]
			if nextS > li {
				continue
			}
			dp[i][j] = Max(dp[i][j], dp[i-1][nextS]+a[i-1])
		}
	}
	var ans int
	for j := 0; j <= li; j++ {
		ans = Max(ans, dp[n][j])
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
