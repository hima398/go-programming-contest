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

func solveHonestly(n int, a, b []int) int {
	const p = 998244353
	mask := 1<<n - 1
	var ans int
	for k := 1; k <= mask; k++ {
		var mx, s int
		for i := 0; i < n; i++ {
			if k>>i&1 > 0 {
				mx = Max(mx, a[i])
				s += b[i]
			}
		}
		if mx >= s {
			ans++
		}
	}
	return ans
}

func solve(n int, a, b []int) int {
	const p = 998244353
	const maxA = 5000
	type node struct {
		a, b int
	}
	var nodes []node
	for i := range a {
		nodes = append(nodes, node{a[i], b[i]})
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].a == nodes[j].a {
			return nodes[i].b < nodes[j].b
		}
		return nodes[i].a < nodes[j].a
	})
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxA+1)
	}
	dp[0][0] = 1
	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j <= maxA; j++ {
			dp[i+1][j] = dp[i][j]
			if nodes[i].b <= j {
				dp[i+1][j] += dp[i][j-nodes[i].b]
				dp[i+1][j] %= p
			}
			if j <= nodes[i].a-nodes[i].b {
				ans += dp[i][j]
				ans %= p
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
	b := nextIntSlice(n)

	ans := solve(n, a, b)
	fmt.Println(ans)
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

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
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
