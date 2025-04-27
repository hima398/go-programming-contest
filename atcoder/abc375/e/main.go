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
	var a, b []int
	for i := 0; i < n; i++ {
		a, b = append(a, nextInt()), append(b, nextInt())
	}

	ans := solve(n, a, b)

	Print(ans)
}

func solve(n int, a, b []int) int {
	var s int
	for _, bi := range b {
		s += bi
	}
	if s%3 != 0 {
		return -1
	}
	//3チームの強さが等しくなる時の値
	s /= 3

	const INF = 1 << 60
	const MaxSize = 501

	//i番目の人まで見てチーム1の強さがj、チーム2の強さがkのとき人数を変更した最大値
	dp := make([][MaxSize][MaxSize]int, n+1)
	//初期化
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = INF
			}
		}
	}

	dp[0][0][0] = 0
	for i := 0; i < n; i++ {
		for j := range dp[i] {
			for k := range dp[i][j] {
				if dp[i][j][k] == INF {
					continue
				}
				switch a[i] {
				case 1:
					next := j + b[i]
					if next < MaxSize {
						dp[i+1][next][k] = Min(dp[i+1][next][k], dp[i][j][k])
					}
					//チーム2に変更
					next = k + b[i]
					if next < MaxSize {
						dp[i+1][j][next] = Min(dp[i+1][j][next], dp[i][j][k]+1)
					}
					//チーム3に変更
					dp[i+1][j][k] = Min(dp[i+1][j][k], dp[i][j][k]+1)
				case 2:
					next := k + b[i]
					if next < MaxSize {
						dp[i+1][j][next] = Min(dp[i+1][j][next], dp[i][j][k])
					}
					//チーム1に変更
					next = j + b[i]
					if next < MaxSize {
						dp[i+1][next][k] = Min(dp[i+1][next][k], dp[i][j][k]+1)
					}
					//チーム3に変更
					dp[i+1][j][k] = Min(dp[i+1][j][k], dp[i][j][k]+1)
				case 3:
					dp[i+1][j][k] = Min(dp[i+1][j][k], dp[i][j][k])
					//チーム1に変更
					next := j + b[i]
					if next < MaxSize {
						dp[i+1][next][k] = Min(dp[i+1][next][k], dp[i][j][k]+1)
					}
					//チーム2に変更
					next = k + b[i]
					if next < MaxSize {
						dp[i+1][j][next] = Min(dp[i+1][j][next], dp[i][j][k]+1)
					}
				}
			}
		}
	}
	if dp[n][s][s] == INF {
		return -1
	}
	return dp[n][s][s]
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
