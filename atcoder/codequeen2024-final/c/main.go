package main

import (
	"bufio"
	"fmt"
	"math"
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

	n, k, t := nextInt(), nextInt(), nextInt()
	p := nextIntSlice(n)

	ans := solve(n, k, t, p)

	Print(ans)
}

func solve(n, k, t int, p []int) int {
	s := make([]int, n+1)
	for i, pi := range p {
		s[i+1] = s[i] + pi
	}

	const INF = 1 << 60

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	for i := range dp {
		dp[i][0] = 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			dp[i+1][j] = Min(dp[i+1][j], dp[i][j])

			next := i + t
			if next > n {
				continue
			}
			dp[next][j+1] = Min(dp[next][j+1], dp[i][j]+s[next]-s[i])
		}
	}
	ans := math.MaxInt
	for i := range dp {
		//fmt.Println(dp[i])
		ans = Min(ans, dp[i][k])
	}

	return ans
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
