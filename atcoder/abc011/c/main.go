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
	ngs := nextIntSlice(3)
	ans := solve(n, ngs)
	PrintString(ans)
}

func solve(n int, ngs []int) string {
	dp := make([][]int, 100+1)
	for i := 0; i <= 100; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = 101
		}
	}
	dp[0][0] = 0
	ngMap := make(map[int]struct{})
	for _, v := range ngs {
		ngMap[v] = struct{}{}
	}
	for i := 1; i <= 100; i++ {
		for j := 0; j <= n; j++ {
			for k := 1; k <= 3; k++ {
				nj := j + k
				if nj > n {
					continue
				}
				if _, found := ngMap[nj]; found {
					continue
				}
				dp[i][nj] = Min(dp[i][nj], dp[i-1][j]+1)
			}
		}
	}
	for i := 1; i <= 100; i++ {
		if 0 < dp[i][n] && dp[i][n] <= 100 {
			return "YES"
		}
	}
	return "NO"
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
