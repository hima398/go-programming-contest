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

	n, x := nextInt(), nextInt()
	var v, a, c []int
	for i := 0; i < n; i++ {
		v = append(v, nextInt()-1)
		a = append(a, nextInt())
		c = append(c, nextInt())
	}
	ans := solve(n, x, v, a, c)
	Print(ans)
}

func solve(n, x int, v, a, c []int) int {
	//ビタミンkについて、i個の食べ物を見てカロリー摂取がjのときビタミンの最大
	dp := make([][][]int, 3)
	for k := range dp {
		dp[k] = make([][]int, n+1)
		for i := range dp[k] {
			dp[k][i] = make([]int, x+1)
		}
	}
	for k := range dp {
		for i := 0; i < n; i++ {
			for j := 0; j <= x; j++ {
				dp[k][i+1][j] = Max(dp[k][i+1][j], dp[k][i][j])
				if k == v[i] {
					nj := j + c[i]
					if nj > x {
						continue
					}
					dp[k][i+1][nj] = Max(dp[k][i+1][nj], dp[k][i][j]+a[i])
				}
			}
		}
	}
	//for k := 0; k < 3; k++ {
	//	fmt.Println(dp[k][n])
	//}
	//カロリーx以下で3つの栄養がy以上摂取できるか？
	check := func(y int) bool {
		for i := 0; i <= x; i++ {
			for j := 0; j <= x-i; j++ {
				k := x - i - j
				if dp[0][n][i] >= y && dp[1][n][j] >= y && dp[2][n][k] >= y {
					return true
				}
			}
		}
		return false
	}
	ok, ng := 0, n*2*int(1e5)+1
	for ng-ok > 1 {
		mid := (ok + ng) / 2
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
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
