package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, s int, a []int) string {
	const maxS = 60 * int(1e4)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxS+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= maxS; j++ {
			dp[i][j] = dp[i-1][j]
			pj := j - a[i-1]
			if pj < 0 {
				continue
			}
			dp[i][j] += dp[i-1][pj]
		}
	}
	if dp[n][s] > 0 {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, s := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, s, a)
	PrintString(ans)
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
