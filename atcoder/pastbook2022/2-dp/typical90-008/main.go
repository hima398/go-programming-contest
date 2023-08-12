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
	ans := solve(n, s)
	PrintInt(ans)
}

func solve(n int, s string) int {
	const p = int(1e9) + 7
	const t = "atcoder"
	//i文字目まで見て、atcoderのj文字目まで作れるパターンの数
	dp := make([][8]int, n+1)
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 8; j++ {
			dp[i+1][j] += dp[i][j]
			dp[i+1][j] %= p
			if j < 7 && s[i] == t[j] {
				dp[i+1][j+1] += dp[i][j]
				dp[i+1][j+1] %= p
			}
		}
	}
	return dp[n][7]
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
