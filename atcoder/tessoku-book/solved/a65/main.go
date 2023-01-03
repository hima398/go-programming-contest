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

	n := nextInt()
	a := nextIntSlice(n - 1)
	ans := solve(n, a)
	PrintHorizonaly(ans)
}

func solve(n int, a []int) []int {
	//0-indexed
	for i := range a {
		a[i]--
	}
	e := make([][]int, n)
	for i, par := range a {
		e[i+1] = append(e[i+1], par)
		e[par] = append(e[par], i+1)
	}
	dp := make([]int, n)
	var dfs func(cur, par int) int
	dfs = func(cur, par int) int {
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			dp[cur] += dfs(next, cur)
		}
		return dp[cur] + 1
	}
	dfs(0, -1)
	return dp
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
