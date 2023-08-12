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

	h, n := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(h, n, a, b)
	PrintInt(ans)
}

func solve(h, n int, a, b []int) int {
	const INF = 1 << 60
	var max int
	for _, ai := range a {
		max = Max(max, ai)
	}
	dp := make([]int, h+max+1)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 1; i <= h+max; i++ {
		for j := 0; j < n; j++ {
			if i-a[j] < 0 {
				continue
			}
			dp[i] = Min(dp[i], dp[i-a[j]]+b[j])
		}
	}

	ans := INF
	for i := h; i <= h+max; i++ {
		ans = Min(ans, dp[i])
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
