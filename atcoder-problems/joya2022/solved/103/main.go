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

	n, m := nextInt(), nextInt()
	a, b, c, d, e, f := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	var x, y []int
	for i := 0; i < m; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(n, m, a, b, c, d, e, f, x, y)
	PrintInt(ans)
}

func solve(n, m, a, b, c, d, e, f int, x, y []int) int {
	const p = 998244353

	obs := make(map[int]map[int]struct{})
	for i := 0; i < m; i++ {
		if obs[x[i]] == nil {
			obs[x[i]] = make(map[int]struct{})
		}
		obs[x[i]][y[i]] = struct{}{}
	}
	//i回の移動のうち、(x+A, y+B)にj回移動、(x+c, y+d)にk回移動する経路の数
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k <= i-j; k++ {
				l := i - j - k
				curX, curY := j*a+k*c+l*e, j*b+k*d+l*f
				if _, found := obs[curX+a][curY+b]; !found {
					dp[i+1][j+1][k] += dp[i][j][k]
					dp[i+1][j+1][k] %= p
				}
				if _, found := obs[curX+c][curY+d]; !found {
					dp[i+1][j][k+1] += dp[i][j][k]
					dp[i+1][j][k+1] %= p
				}
				if _, found := obs[curX+e][curY+f]; !found {
					dp[i+1][j][k] += dp[i][j][k]
					dp[i+1][j][k] %= p
				}
			}
		}
	}
	var ans int
	for j := 0; j <= n; j++ {
		for k := 0; k <= n-j; k++ {
			ans += dp[n][j][k]
			ans %= p
		}
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
