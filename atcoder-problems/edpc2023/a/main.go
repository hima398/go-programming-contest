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

	n := nextInt()
	h := nextIntSlice(n)

	ans := solve(n, h)

	Print(ans)
}

func solve(n int, h []int) int {
	dp := make([]int, n)
	for i := range dp {
		//INF > 10**5 * 10**4
		dp[i] = int(1e9) + 1
	}
	dp[0] = 0
	dp[1] = Abs(h[1] - h[0])
	for i := 2; i < n; i++ {
		dp[i] = Min(dp[i-1]+Abs(h[i]-h[i-1]), dp[i-2]+Abs(h[i]-h[i-2]))
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

func PrintInt(x int) {
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
