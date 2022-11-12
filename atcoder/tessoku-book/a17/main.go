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

const INF = 1 << 60

func solve(n int, a, b []int) []int {
	a = append([]int{0}, a...)
	b = append([]int{0, 0}, b...)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = INF
	}
	dp[1] = 0
	dp[2] = a[1]
	for i := 3; i <= n; i++ {
		dp[i] = Min(dp[i-1]+a[i-1], dp[i-2]+b[i-1])
	}

	var ans []int
	cur := n
	for {
		ans = append(ans, cur)
		if dp[cur-1]+a[cur-1] == dp[cur] {
			cur--
		} else {
			cur -= 2
		}
		if cur <= 0 {
			break
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n - 1)
	b := nextIntSlice(n - 2)
	ans := solve(n, a, b)
	PrintInt(len(ans))
	PrintHorizonaly(ans)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
