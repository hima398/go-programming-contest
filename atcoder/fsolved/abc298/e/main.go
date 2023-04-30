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
	a, b, p, q := nextInt(), nextInt(), nextInt(), nextInt()
	ans := solve(n, a, b, p, q)
	PrintInt(ans)
}

func solve(n, a, b, p, q int) int {
	const mod = 998244353
	//高橋君が地点i、青木君が地点jにいて、
	//ターンがk(k=0高橋君、k=1青木君)のとき、高橋君が勝つ確率
	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, n+1)
	}
	for j := 0; j < n; j++ {
		for k := 0; k < 2; k++ {
			dp[n][j][k] = 1
		}
	}
	ip, iq := Inv(p, mod), Inv(q, mod)
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			for k := 1; k <= p; k++ {
				dp[i][j][0] += dp[Min(n, i+k)][j][1]
			}
			dp[i][j][0] *= ip
			dp[i][j][0] %= mod
			for k := 1; k <= q; k++ {
				dp[i][j][1] += dp[i][Min(n, j+k)][0]
			}
			dp[i][j][1] *= iq
			dp[i][j][1] %= mod
		}
	}

	return dp[a][b][0]
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
