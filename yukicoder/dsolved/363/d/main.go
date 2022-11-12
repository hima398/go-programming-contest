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

	n, m := nextInt(), nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(m)
	}
	ans := solve(n, m, a)
	PrintInt(ans)
}

func solve(n, m int, a [][]int) int {
	const INF = 1 << 60
	if n == 1 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for j := 0; j < m; j++ {
		dp[0][j] = a[0][j]
	}
	min := INF
	for j := 0; j < m; j++ {
		dp[1][j] = dp[0][j] + a[1][j]
		min = Min(min, dp[1][j])
	}
	for i := 2; i < n; i++ {
		for j := 0; j < m; j++ {
			if dp[i-1][j] == min {
				dp[i][j] = min + a[i][j]
			} else {
				dp[i][j] = Min(dp[i-1][j]+a[i][j], min+a[i][j]+a[i-1][j])
			}
		}
		min = INF
		for j := 0; j < m; j++ {
			min = Min(min, dp[i][j])
		}
	}

	ans := 1 << 60
	for j := 0; j < m; j++ {
		ans = Min(ans, dp[n-1][j])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
