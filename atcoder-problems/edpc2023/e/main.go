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

	n, lw := nextInt(), nextInt()
	var w, v []int
	for i := 0; i < n; i++ {
		w = append(w, nextInt())
		v = append(v, nextInt())
	}

	ans := solve(n, lw, w, v)

	Print(ans)
}

func solve(n, lw int, w, v []int) int {
	const INF = 1 << 60
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, int(1e5)+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	dp[0][v[0]] = w[0]
	for i := 1; i < n; i++ {
		for j := 0; j <= int(1e5); j++ {
			dp[i][j] = dp[i-1][j]
			pj := j - v[i]
			if pj < 0 {
				continue
			}
			dp[i][j] = Min(dp[i][j], dp[i-1][pj]+w[i])
		}
	}
	var ans int
	for i, v := range dp[n-1] {
		if v > 0 && v <= lw {
			ans = Max(ans, i)
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
