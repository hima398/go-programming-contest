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

	n, x, y := nextInt(), nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(n, x, y, a, b)
	Print(ans)
}

func solve(n, x, y int, a, b []int) int {
	const INF = 1 << 60
	dp := make([][]int, n+1)
	for j := range dp {
		dp[j] = make([]int, x+1)
		for k := range dp[j] {
			dp[j][k] = INF
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		next := make([][]int, n+1)
		for j := range dp {
			next[j] = make([]int, x+1)
			for k := range dp[j] {
				next[j][k] = INF
			}
		}

		for j := 0; j <= i; j++ {
			for k := 0; k <= x; k++ {
				//i番目の料理を食べない
				next[j][k] = Min(next[j][k], dp[j][k])
				//i番目の料理を食べる
				nj, nk := j+1, k+a[i]
				if nk <= x {
					next[nj][nk] = Min(next[nj][nk], dp[j][k]+b[i])
				}
			}
		}
		dp = next
	}
	var ans int
	for j := 0; j <= n; j++ {
		for k := 0; k <= x; k++ {
			if dp[j][k] <= y {
				ans = Max(ans, j+1)
			}
		}
	}
	return Min(ans, n)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
