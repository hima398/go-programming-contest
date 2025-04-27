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
	a := nextIntSlice(n)

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a []int) int {
	//i番目のモンスターに出会って、モンスターを倒す数のmodがjの時に得られる経験値
	dp := make([][2]int, n+1)
	dp[0][1] = -1
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			if dp[i][j] < 0 {
				continue
			}
			//モンスターを逃す
			dp[i+1][j] = Max(dp[i+1][j], dp[i][j])
			//モンスターを倒す
			if j == 0 {
				dp[i+1][1] = Max(dp[i+1][1], dp[i][j]+a[i])
			} else {
				dp[i+1][0] = Max(dp[i+1][0], dp[i][j]+2*a[i])
			}
		}
	}
	ans := Max(dp[n][0], dp[n][1])
	return ans
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
