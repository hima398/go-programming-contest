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

	n, m, k := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, m, k, a)
	PrintInt(ans)
}

func solve(n, m, k int, a []int) int {
	const INF = 1 << 60
	//コストの前計算
	//iから長さjまで[i, i+j)のオレンジを箱詰めするコスト
	cost := make([][]int, n+1)
	for i := range cost {
		cost[i] = make([]int, m+1)
		for j := range cost[i] {
			cost[i][j] = INF
		}
	}
	for i := range cost {
		min, max := INF, 0
		for j := 1; j <= m; j++ {
			if i+j-1 >= n {
				break
			}
			idx := i + j - 1
			min = Min(min, a[idx])
			max = Max(max, a[idx])
			cost[i][j] = k + (max-min)*j
		}
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := Max(i-m, 0); j < i; j++ {
			l := i - j
			dp[i] = Min(dp[i], dp[j]+cost[j][l])
		}
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
