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
	var s []string
	var c []int
	for i := 0; i < m; i++ {
		s = append(s, nextString())
		c = append(c, nextInt())
	}
	ans := solve(n, m, s, c)
	PrintInt(ans)
}

func solve(n, m int, s []string, c []int) int {
	const INF = 1 << 60
	var patterns []int
	for _, si := range s {
		p := 0
		for idx, r := range si {
			if r == 'Y' {
				p |= 1 << idx
			}
		}
		patterns = append(patterns, p)
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		for pat := 0; pat < 1<<n; pat++ {
			//i番目の部品を買わない
			dp[i][pat] = Min(dp[i][pat], dp[i-1][pat])

			//i番目の部品を買う
			next := pat | patterns[i-1]
			dp[i][next] = Min(dp[i][next], dp[i-1][pat]+c[i-1])
		}
	}

	ans := dp[m][(1<<n)-1]
	if ans == INF {
		return -1
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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
