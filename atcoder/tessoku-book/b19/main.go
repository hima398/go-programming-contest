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
	const INF = 1 << 60
	var maxV int
	for _, vi := range v {
		maxV += vi
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxV+1)
		for j := 0; j <= maxV; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= maxV; j++ {
			//i番目を持ち帰らない
			dp[i][j] = dp[i-1][j]
			//i番目を持ち帰る
			pj := j - v[i-1]
			if pj < 0 {
				continue
			}
			dp[i][j] = Min(dp[i][j], dp[i-1][pj]+w[i-1])
		}
	}
	var ans int
	for j := 0; j <= maxV; j++ {
		if dp[n][j] <= maxW {
			ans = j
		}
	}
	//fmt.Println(dp[n])
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

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
