package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	ans := solve(s)
	PrintInt(ans)
}

func solve(s string) int {
	const INF = 1 << 60

	s = "0" + s
	n := len(s)
	//上の位からi桁目を見て繰り下がりの状態がjの時の紙幣の最小枚数
	//j=0：繰り下がりなし、j=1：繰り下がりあり
	dp := make([][2]int, n+1)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[n][0] = 0
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < 2; j++ {
			v := int(s[i] - '0')
			v += j
			if 0 < v {
				dp[i][1] = Min(dp[i][1], dp[i+1][j]+(10-v))
			}
			if v < 10 {
				dp[i][0] = Min(dp[i][0], dp[i+1][j]+v)
			}
		}
	}
	return dp[0][0]
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
