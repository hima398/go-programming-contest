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

	n, lw := nextInt(), nextInt()
	var w, v []int
	for i := 0; i < n; i++ {
		w, v = append(w, nextInt()), append(v, nextInt())
	}

	ans := solve(n, lw, w, v)
	//ans := solveHonestly(n, lw, w, v)

	Print(ans)
}

func solve(n, lw int, w, v []int) int {
	const INF = 1 << 60
	s := make([][]int, lw+1)
	for i := range s {
		s[i] = make([]int, lw+1)
		for j := range s[i] {
			s[i][j] = -INF
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= lw; j++ {
			fmt.Println(i, j)
			s[w[i]][j] = Max(s[w[i]][j], j*v[i]-j*j)
		}
	}
	fmt.Println(s)

	dp := make([][]int, lw+1)
	for i := range dp {
		dp[i] = make([]int, lw+1)
		for j := range dp[i] {
			dp[i][j] = -INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= lw; i++ {
		for k := 0; k <= lw; k++ {
			dp[i][k] = dp[i-1][k]
		}
		for j := 0; j <= Ceil(lw, i); j++ {
			for k := 0; k <= lw; k++ {
				if k-i < 0 {
					continue
				}
				dp[i][k] = Max(dp[i][k], dp[i-1][k-i]+s[i][j])
			}
		}
	}
	fmt.Println(dp)

	ans := -INF
	for k := 0; k <= lw; k++ {
		ans = Max(ans, dp[lw][k])
	}

	return ans
}

func solveHonestly(n, lw int, w, v []int) int {
	const INF = 1 << 60
	//i番目の品物まで見てj個選んで重さがkになった時の嬉しさの最大値
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, lw+1)
		for k := range dp[i] {
			dp[i][k] = -INF
		}
	}
	dp[0][0] = 0
	//fmt.Println(dp)
	for i := 0; i < n; i++ {
		for k := range dp[i] {
			dp[i+1][k] = Max(dp[i+1][k], dp[i][k])
		}
		for j := 0; j <= Ceil(lw, w[i]); j++ {
			for k := range dp[i] {
				//存在しない状態
				if dp[i][k] < 0 {
					continue
				}
				nextK := k + j*w[i]
				//容量を超えてしまう
				if nextK > lw {
					continue
				}
				dp[i+1][nextK] = Max(dp[i+1][nextK], dp[i][k]+j*v[i]-j*j)
			}
		}
	}
	var ans int
	for k := range dp[n] {
		ans = Max(ans, dp[n][k])
	}
	return ans
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
