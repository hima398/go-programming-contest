package main

import (
	"bufio"
	"fmt"
	"math"
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
	s := nextString()
	c := nextIntSlice(n)
	d := nextIntSlice(n)

	ans := solve(n, s, c, d)
	PrintInt(ans)
}

func solve(n int, s string, c, d []int) int {
	const INF = math.MaxInt64
	//i番目の文字列まで見て、j個のカッコが開いている時の最小コスト
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			//i番目の文字列を削除する
			dp[i+1][j] = Min(dp[i+1][j], dp[i][j]+d[i])
			//i番目の文字をそのまま使うか、反転して'('にする
			if s[i] == '(' {
				dp[i+1][j+1] = Min(dp[i+1][j+1], dp[i][j])
			} else {
				dp[i+1][j+1] = Min(dp[i+1][j+1], dp[i][j]+c[i])
			}
			//i番目の文字をそのまま使うか、反転して')'にする
			//ただし、前に1つ以上'('が必要
			if j > 0 {
				if s[i] == '(' {
					dp[i+1][j-1] = Min(dp[i+1][j-1], dp[i][j]+c[i])
				} else {
					dp[i+1][j-1] = Min(dp[i+1][j-1], dp[i][j])
				}
			}
		}
	}
	return dp[n][0]
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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
