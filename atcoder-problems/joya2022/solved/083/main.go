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

	n, k, d := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := solve(n, k, d, a)
	PrintInt(ans)
}

func solve(n, k, d int, a []int) int {
	dp := make([][][]int, n+5)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, k+5)
		for j := range dp[i] {
			dp[i][j] = make([]int, d+5)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	dp[0][0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= Min(i, k); j++ {
			for ii := 0; ii < d; ii++ {
				dp[i+1][j][ii] = Max(dp[i+1][j][ii], dp[i][j][ii])
			}
			for ii := 0; ii < d; ii++ {
				if dp[i][j][ii] == -1 {
					continue
				}
				next := (ii + a[i]) % d
				dp[i+1][j+1][next] = Max(dp[i+1][j+1][next], dp[i][j][ii]+a[i])
			}
		}
	}

	return dp[n][k][0]
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
