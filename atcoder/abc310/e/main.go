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
	//i番目の文字を見て0 or 1になる数
	dp := make([][2]int, n)
	var ans int
	if s[0] == '0' {
		dp[0][0] = 1
	} else {
		dp[0][1] = 1
		ans++
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			dp[i][0] = 1
			dp[i][1] = dp[i-1][0] + dp[i-1][1]
		} else {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
		ans += dp[i][1]
	}
	return ans
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
