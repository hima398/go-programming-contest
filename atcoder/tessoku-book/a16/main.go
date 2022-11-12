package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

const INF = 1 << 60

func solve(n int, a, b []int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = INF
	}
	dp[1] = 0
	dp[2] = a[1]
	for i := 3; i <= n; i++ {
		dp[i] = Min(dp[i-1]+a[i-1], dp[i-2]+b[i-1])
	}
	return dp[n]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = INF
		b[i] = INF
	}
	for i := 1; i < n; i++ {
		a[i] = nextInt()
	}
	for i := 2; i < n; i++ {
		b[i] = nextInt()
	}
	ans := solve(n, a, b)
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
	if x < y {
		return x
	}
	return y
}
