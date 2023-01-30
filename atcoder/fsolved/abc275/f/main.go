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
	a := nextIntSlice(n)
	ans := solve(n, m, a)
	PrintVertically(ans)
}

func solve(n, m int, a []int) []int {
	const INF = 1 << 60
	//a_iまで見て合計がjになり、k=0ならば削除k=1ならばa_iを採用する場合の操作回数の最小値
	dp := make([][][2]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][2]int, m+1)
		for j := 0; j <= m; j++ {
			for k := 0; k < 2; k++ {
				dp[i][j][k] = INF
			}
		}
	}
	dp[0][0][1] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			dp[i+1][j][0] = Min(dp[i][j][0], dp[i][j][1]+1)
			nextJ := j + a[i]
			if nextJ > m {
				continue
			}
			dp[i+1][nextJ][1] = Min(dp[i][j][0], dp[i][j][1])
		}
	}
	var ans []int
	for j := 1; j <= m; j++ {
		v := Min(dp[n][j][0], dp[n][j][1])
		if v == INF {
			v = -1
		}
		ans = append(ans, v)
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
