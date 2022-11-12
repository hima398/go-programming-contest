package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a, b []int) int {
	const INF = 1 << 60
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -INF
	}
	dp[0] = 0
	for i := 0; i < n-1; i++ {
		dp[a[i]-1] = Max(dp[a[i]-1], dp[i]+100)
		dp[b[i]-1] = Max(dp[b[i]-1], dp[i]+150)
	}
	///fmt.Println(dp)
	return dp[n-1]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n - 1)
	b := nextIntSlice(n - 1)
	ans := solve(n, a, b)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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
