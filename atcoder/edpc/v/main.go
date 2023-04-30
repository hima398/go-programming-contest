package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, m int, x, y []int) []int {
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[x[i]] = append(e[x[i]], y[i])
		e[y[i]] = append(e[y[i]], x[i])
	}
	memo := make([]int, n)
	var dfs1 func(cur, par int)
	dfs1 = func(cur, par int) {
		memo[cur] = 1
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			dfs1(next, cur)
			memo[cur] *= memo[next] + 1
			memo[cur] %= m
		}
	}
	dfs1(0, -1)

	ans := make([]int, n)
	var dfs2 func(cur, par int)
	dfs2 = func(cur, par int) {
		ans[cur] = 1
		for _, next := range e[cur] {
			ans[cur] *= memo[next] + 1
			ans[cur] %= m
		}
		nc := len(e[cur])
		l, r := make([]int, nc), make([]int, nc)
		for i, next := range e[cur] {
			l[i] = memo[next] + 1
			r[i] = memo[next] + 1
		}
		for i := 1; i < nc; i++ {
			l[i] *= l[i-1]
			l[i] %= m
		}
		for i := nc - 2; i >= 0; i-- {
			r[i] *= r[i+1]
			r[i] %= m
		}
		for i, next := range e[cur] {
			if next == par {
				continue
			}
			memo[cur] = 1
			if i > 0 {
				memo[cur] *= l[i-1]
				memo[cur] %= m
			}
			if i+1 < nc {
				memo[cur] *= r[i+1]
				memo[cur] %= m
			}
			dfs2(next, cur)
		}
	}
	dfs2(0, -1)

	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var x, y []int
	for i := 0; i < n-1; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt()-1)
	}
	ans := solve(n, m, x, y)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
