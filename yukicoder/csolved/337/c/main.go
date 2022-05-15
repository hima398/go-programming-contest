package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n, x, y int) int {
	const p = 998244353
	if x > y {
		x, y = y, x
	}
	d := y - x
	if y != n {
		d--
	}
	if x != 1 {
		d--
	}
	if d < 0 {
		return 0
	}
	dp := make([]int, d+1)
	dp[0] = 1
	for i := 0; i < d; i++ {
		dp[i+1] += dp[i]
		dp[i+1] %= p
		if i+3 <= d {
			dp[i+3] += dp[i]
			dp[i+3] %= p
		}
	}
	return dp[d]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x, y := nextInt(), nextInt(), nextInt()
	ans := solve(n, x, y)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
