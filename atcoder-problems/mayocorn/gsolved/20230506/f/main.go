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

	n, s := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, s, a)
	PrintInt(ans)
}

func solve(n, s int, a []int) int {
	const p = 998244353

	//i番目まで見て、tの部分集合の合計がsになる個数の和
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, s+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= s; j++ {
			dp[i+1][j] += 2 * dp[i][j]
			dp[i+1][j] %= p
			if j+a[i] > s {
				continue
			}
			dp[i+1][j+a[i]] += dp[i][j]
			dp[i+1][j+a[i]] %= p
		}
	}

	return dp[n][s]
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

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
