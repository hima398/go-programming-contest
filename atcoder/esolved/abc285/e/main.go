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
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintInt(ans)
}

func solve(n int, a []int) int {
	//休日同士の間隔がi日である部分の生産量の総和
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[(i-1)/2]
	}
	//i日目まで見て、i日目が休みの時の生産量
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		mx := 0
		for j := 0; j < i; j++ {
			mx = Max(mx, dp[j]+s[i-j-1])
		}
		dp[i] = mx
	}
	return dp[n]
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
