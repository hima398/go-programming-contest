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

	n, p := nextInt(), nextInt()
	ans := solve(n, p)
	PrintHorizonaly(ans)
}

func solve(n, p int) []int {
	//i番目のコまで見て、j本の辺を削除して連結の状態がk(k=0：非連結、k=1：連結)
	dp := make([][][2]int, n+2)
	for i := range dp {
		dp[i] = make([][2]int, n+2)
	}
	dp[0][0][1] = 1
	dp[0][1][0] = 1

	for i := 0; i < n; i++ {
		ni := i + 1
		for j := 0; j < n; j++ {
			//i番目に1本も変を切らなければi-1番目のどちらのパターンも連結になる
			dp[ni][j][1] += dp[i][j][0] + dp[i][j][1]
			dp[ni][j][1] %= p

			//i番目にコの縦の辺を削除
			dp[ni][j+1][0] += dp[i][j][0]
			dp[ni][j+1][0] %= p
			dp[ni][j+1][1] += dp[i][j][1]
			dp[ni][j+1][1] %= p

			//i番目にコの横の辺を削除
			dp[ni][j+1][1] += 2 * dp[i][j][1]
			dp[ni][j+1][1] %= p

			//i番目のコから2本の辺を削除してi+1番目のつなぎ方によって連結にできるもの
			dp[ni][j+2][0] += 2 * dp[i][j][1]
			dp[ni][j+2][0] %= p
		}
	}
	var ans []int
	for j := 1; j < n; j++ {
		ans = append(ans, dp[n-1][j][1])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
