package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, k int) int {
	const p = 998244353

	//jの位置にいる時、残りi回でNに辿り着ける確率
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= k; i++ {
		dp[i][n] = 1
	}
	for i := 0; i < k; i++ {
		var s, inv int
		for j := n - 1; j >= 0; j-- {
			s += dp[i][j+1]
			s %= p
			inv++
			inv %= p

			dp[i+1][j] = s * Inv(inv, p) % p
		}
	}
	return dp[k][0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	ans := solve(n, k)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
