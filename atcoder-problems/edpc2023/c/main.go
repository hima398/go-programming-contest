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
	var a, b, c []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
		c = append(c, nextInt())
	}

	ans := solve(n, a, b, c)

	Print(ans)
}

func solve(n int, a, b, c []int) int {
	dp := make([][3]int, n)
	dp[0][0] = a[0]
	dp[0][1] = b[0]
	dp[0][2] = c[0]
	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			switch j {
			case 0:
				dp[i][j] = Max(dp[i-1][1]+a[i], dp[i-1][2]+a[i])
			case 1:
				dp[i][j] = Max(dp[i-1][0]+b[i], dp[i-1][2]+b[i])
			case 2:
				dp[i][j] = Max(dp[i-1][0]+c[i], dp[i-1][1]+c[i])
			}
		}
	}
	var ans int
	for j := range dp[n-1] {
		ans = Max(ans, dp[n-1][j])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
