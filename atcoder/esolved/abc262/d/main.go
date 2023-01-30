package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a []int) (ans int) {
	const p = 998244353

	for m := 1; m <= n; m++ {

		var dp [101][101][101]int
		//dp := make([][][]int, n+1)
		//for i := 0; i <= n; i++ {
		//	dp[i] = make([][]int, n+1)
		//	for j := 0; j <= n; j++ {
		//		dp[i][j] = make([]int, n+1)
		//	}
		//}
		dp[0][0][0] = 1
		for i := 1; i <= n; i++ {
			//j番目を取らない
			for j := 0; j <= m; j++ {
				for k := 0; k <= m; k++ {
					dp[i][j][k] += dp[i-1][j][k]
					dp[i][j][k] %= p
					if j > 0 {
						dp[i][j][(k+a[i-1])%m] += dp[i-1][j-1][k]
						dp[i][j][(k+a[i-1])%m] %= p
					}
				}
			}
			//j番目をとる
			//for j := 1; j <= m; j++ {
			//	for k := 0; k <= m; k++ {
			//		dp[i][j][(k+a[i-1])%m] += dp[i-1][j-1][k]
			//		dp[i][j][(k+a[i-1])%m] %= p
			//	}
			//}
		}
		ans += dp[n][m][0]
		ans %= p
	}

	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	ans := solve(n, a)
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
