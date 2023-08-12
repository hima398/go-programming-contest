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

	n, s := nextInt(), nextInt()
	var w, v []int
	for i := 0; i < n; i++ {
		w = append(w, nextInt())
		v = append(v, nextInt())
	}
	ans := solve(n, s, w, v)
	PrintInt(ans)
}

func solve(n, s int, w, v []int) int {
	//i個目の品物まで見て、重さj詰めた場合の価値の最大値
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, s+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= s; j++ {
			//i番目の荷物を詰めない
			dp[i+1][j] = Max(dp[i+1][j], dp[i][j])

			//i番目の荷物を詰める
			nj := j + w[i]
			//i番目の荷物を入れると容量を超えてしまう
			if nj > s {
				continue
			}
			dp[i+1][nj] = Max(dp[i+1][nj], dp[i][j]+v[i])
		}
	}
	var ans int
	for j := 0; j <= s; j++ {
		ans = Max(ans, dp[n][j])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
