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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k, x := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, k, x, a)
	PrintInt(ans)
}

func solve(n, k, x int, a []int) int {
	//i日目にサポータか否か(0:非サポーター、1:サポーター)で払う金額の最小値
	dp := make([][2]int, n+1)
	dp[0][1] = 1 << 60
	for i := 1; i <= n; i++ {
		dp[i][0] = Min(dp[i-1][0]+a[i-1], dp[i-1][1]+a[i-1])
		dp[i][1] = Min(dp[i-1][0]+x+k, dp[i-1][1]+k)
	}
	//fmt.Println(dp)
	return Min(dp[n][0], dp[n][1])
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
