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

	n := nextInt()
	t := nextIntSlice(n)
	ans := solve(n, t)
	PrintInt(ans)
}

func solve(n int, t []int) int {
	s := 0
	for _, ti := range t {
		s += ti
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, s+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			dp[i][j] = dp[i][j] || dp[i-1][j]
			if j-t[i-1] < 0 {
				continue
			}
			dp[i][j] = dp[i][j] || dp[i-1][j-t[i-1]]
		}
	}
	ans := s
	for j := 0; j <= s; j++ {
		if dp[n][j] {
			ans = Min(ans, Max(j, s-j))
		}
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
