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

	n, k := nextInt(), nextInt()
	h := nextIntSlice(n)

	ans := solve(n, k, h)

	Print(ans)
}

func solve(n, k int, h []int) int {
	const INF = int(1e9) + 1
	dp := make([]int, n)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := Max(i-k, 0); j < i; j++ {
			dp[i] = Min(dp[i], dp[j]+Abs(h[i]-h[j]))
		}
	}
	return dp[n-1]
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
