package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func computeDigits(x int) int {
	if x == 0 {
		return 0
	}
	res := 0
	for x > 0 {
		res++
		x /= 10
	}
	return res
}
func solve(n int) int {
	const p = 998244353
	// fa <= 200+199
	fa := n + (n - 1)
	w := make([]int, fa+1)
	for i := 1; i <= fa; i++ {
		w[i-computeDigits(i)+1]++
	}
	//fmt.Println(w)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, fa+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= fa; j++ {
			for k := 1; k <= fa; k++ {
				nj := j + k
				if nj <= fa {
					dp[i+1][nj] += dp[i][j] * w[k]
					dp[i+1][nj] %= p
				}
			}
		}
	}
	//fmt.Println(dp)
	return dp[n][fa]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ans := solve(n)
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
