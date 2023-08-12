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
	s := nextString()
	c := nextIntSlice(n)
	d := nextIntSlice(n)

	ans := solve(n, s, c, d)

	PrintInt(ans)
}

func solve(n int, s string, c, d []int) int {
	const INF = 1 << 60
	//i番目の文字まで見て、カッコがj個開いている状態を作るコストの最小値
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			switch s[i-1] {
			case '(':
				//そのまま採用すると開くカッコが1つ増える
				if j > 0 {
					dp[i][j] = Min(dp[i][j], dp[i-1][j-1])
				}
				//反転させると開くカッコが1つ減る
				if j < n {
					dp[i][j] = Min(dp[i][j], dp[i-1][j+1]+c[i-1])
				}
				//i文字目を使わない(削除)
				dp[i][j] = Min(dp[i][j], dp[i-1][j]+d[i-1])
			case ')':
				//そのまま採用するとカッコが1つ減る
				if j < n {
					dp[i][j] = Min(dp[i][j], dp[i-1][j+1])
				}
				//反転させると開くカッコが1つ増える
				if j > 0 {
					dp[i][j] = Min(dp[i][j], dp[i-1][j-1]+c[i-1])
				}
				//i文字目を使わない(削除)
				dp[i][j] = Min(dp[i][j], dp[i-1][j]+d[i-1])
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
