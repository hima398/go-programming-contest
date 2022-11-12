package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, maxW int, w, v []int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxW+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= maxW; j++ {
			//i番目を持ち帰らない
			dp[i][j] = dp[i-1][j]
			//i番目を持ち帰る
			pj := j - w[i-1]
			if pj < 0 {
				continue
			}
			dp[i][j] = Max(dp[i][j], dp[i-1][pj]+v[i-1])
		}
	}
	var ans int
	for j := 0; j <= maxW; j++ {
		ans = Max(ans, dp[n][j])
	}
	return ans
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, maxW := nextInt(), nextInt()
	var w, v []int
	for i := 0; i < n; i++ {
		w = append(w, nextInt())
		v = append(v, nextInt())
	}
	ans := solve(n, maxW, w, v)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
