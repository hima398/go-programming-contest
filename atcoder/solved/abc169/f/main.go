package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n, s int, a []int) int {
	const p = 998244353

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, s+1)
	}
	dp[0][0] = Pow(2, n, p)
	inv := Inv(2, p)
	for i := 0; i < n; i++ {
		for j := 0; j <= s; j++ {
			dp[i+1][j] += dp[i][j]
			dp[i+1][j] %= p
			if j+a[i] <= s {
				dp[i+1][j+a[i]] += dp[i][j] * inv
				dp[i+1][j+a[i]] %= p
			}
		}
	}
	return dp[n][s]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, s := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := solve(n, s, a)

	fmt.Println(ans)
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
