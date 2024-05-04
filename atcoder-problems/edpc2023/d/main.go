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
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, lw+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	dp[0][w[0]] = v[0]
	for i := 1; i < n; i++ {
		for j := 0; j <= lw; j++ {
			dp[i][j] = dp[i-1][j]
			pj := j - w[i]
			if pj < 0 {
				continue
			}
			if dp[i-1][pj] < 0 {
				continue
			}
			dp[i][j] = Max(dp[i][j], dp[i-1][pj]+v[i])
		}
	}

	var ans int
	for _, v := range dp[n-1] {
		ans = Max(ans, v)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
