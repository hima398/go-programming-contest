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
	a := nextIntSlice(n)
	b := nextIntSlice(n)

	ok := solve(n, k, a, b)

	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(n, k int, a, b []int) bool {
	dp := make([][2]bool, n)
	dp[0][0] = true
	dp[0][1] = true
	for i := 1; i < n; i++ {
		dp[i][0] = (dp[i-1][0] && Abs(a[i-1]-a[i]) <= k) || (dp[i-1][1] && Abs(b[i-1]-a[i]) <= k)
		dp[i][1] = (dp[i-1][0] && Abs(a[i-1]-b[i]) <= k) || (dp[i-1][1] && Abs(b[i-1]-b[i]) <= k)
	}
	return dp[n-1][0] || dp[n-1][1]
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
