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

	n, m := nextInt(), nextInt()
	a := nextIntSlice(n)
	c := nextIntSlice(n)
	x := nextIntSlice(m)
	ans := solve(n, m, a, c, x)
	PrintInt(ans)
}

func solve(n, m int, a, c, x []int) int {
	const INF = 1 << 60
	need := make([]bool, n)

	for _, v := range x {
		need[v-1] = true
	}

	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cost[i][j] = INF
		}
	}
	for i := 0; i < n; i++ {
		cur := INF
		for j := i; j < n; j++ {
			cur = Min(cur, c[j])
			cost[i][j] = cur
		}
	}

	//i番目の商品まで見て、そのうちj個の商品を買った時の合計費用
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			//欲しい商品でなければ、i番目を買わない選択肢が取れる
			if !need[i] {
				//i番目の商品を購入しない
				dp[i+1][j] = Min(dp[i+1][j], dp[i][j])
			}
			dp[i+1][j+1] = Min(dp[i+1][j+1], dp[i][j]+a[i]+cost[i-j][i])
		}
	}
	ans := INF
	//fmt.Println(dp[n])
	for j := len(x); j <= n; j++ {
		ans = Min(ans, dp[n][j])
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
