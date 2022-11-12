package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, m int, a [][]int) int {
	const INF = 1 << 60
	b := make([]int, m)
	for i := 0; i < m; i++ {
		pat := 0
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				pat |= 1 << j
			}
		}
		b[i] = pat
	}

	dp := make([][]int, m+1)
	mask := 1<<n - 1
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, mask+1)
		for pat := 0; pat <= mask; pat++ {
			dp[i][pat] = INF
		}
	}
	dp[0][0] = 0
	for i := 0; i < m; i++ {
		//i番目のクーポンを使わない
		for pat := 0; pat <= mask; pat++ {
			dp[i+1][pat] = dp[i][pat]
		}
		//i番目のクーポンを使う
		for pat := 0; pat <= mask; pat++ {
			nextPat := pat | b[i]
			dp[i+1][nextPat] = Min(dp[i+1][nextPat], dp[i][pat]+1)
		}
	}

	if dp[m][mask] == INF {
		return -1
	} else {
		return dp[m][mask]
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = nextIntSlice(n)
	}
	ans := solve(n, m, a)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
