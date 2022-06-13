package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a []int) int {

	sort.Ints(a)
	s := 0
	for _, ai := range a {
		s += ai
	}
	//var dp [mxn + 1][mxs + 1][mxn + 1]int
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, s+1)
		for j := 0; j <= s; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	dp[0][0][0] = 1
	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j <= s; j++ {
			for k := 0; k < n; k++ {
				dp[i+1][j][k] += dp[i][j][k]
				if j+a[i] <= s {
					dp[i+1][j+a[i]][k+1] += dp[i][j][k]
				}
				if k >= 1 && (j+a[i])%k == 0 && j+a[i] >= k*a[i] {
					ans += dp[i][j][k]
				}
			}
		}
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
