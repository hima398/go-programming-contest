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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, k := nextInt(), nextInt(), nextInt()
	ans := solve(n, m, k)
	PrintInt(ans)
}

func solve(n, m, k int) int {
	const p = 998244353
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	im := Inv(m, p)
	//今jにいてルーレットでxが出た時、次の位置
	computeNext := func(j, x int) int {
		if j+x > n {
			over := j + x - n
			return n - over
		} else {
			return j + x
		}
	}
	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			for x := 1; x <= m; x++ {
				next := computeNext(j, x)
				dp[i+1][next] += dp[i][j] * im
				dp[i+1][next] %= p
			}
		}
	}

	var ans int
	for i := 1; i <= k; i++ {
		ans += dp[i][n]
		ans %= p
	}
	return ans
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
