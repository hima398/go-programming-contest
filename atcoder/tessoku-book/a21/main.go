package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, p, a []int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for len := n - 2; len >= 0; len-- {
		for l := 0; l < n-len; l++ {
			r := l + len
			//fmt.Println("l, r = ", l, r)
			var scoreL, scoreR int
			if l > 0 && l <= p[l-1] && p[l-1] <= r {
				scoreL += a[l-1]
			}
			if r < n-1 && l <= p[r+1] && p[r+1] <= r {
				scoreR += a[r+1]
			}
			if l == 0 {
				dp[l][r] = dp[l][r+1] + scoreR
			} else if r == n-1 {
				dp[l][r] = dp[l-1][r] + scoreL
			} else {
				dp[l][r] = Max(dp[l-1][r]+scoreL, dp[l][r+1]+scoreR)
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		ans = Max(ans, dp[i][i])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var p, a []int
	for i := 0; i < n; i++ {
		p = append(p, nextInt()-1)
		a = append(a, nextInt())
	}
	ans := solve(n, p, a)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
