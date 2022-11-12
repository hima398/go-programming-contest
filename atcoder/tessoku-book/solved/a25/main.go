package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(h, w int, c []string) int {
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
	}
	dp[0][0] = 1
	for i := 1; i < h; i++ {
		if c[i][0] == '.' {
			dp[i][0] += dp[i-1][0]
		}
	}
	for j := 1; j < w; j++ {
		if c[0][j] == '.' {
			dp[0][j] += dp[0][j-1]
		}
	}
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if c[i][j] == '.' {
				dp[i][j] += dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[h-1][w-1]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	c := make([]string, h)
	for i := 0; i < h; i++ {
		c[i] = nextString()
	}
	ans := solve(h, w, c)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
