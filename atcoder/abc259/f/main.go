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

func solve(n int, d, u, v, w []int) int {
	const INF = 1 << 60
	type edge struct {
		to, w int
	}
	//辺の本数
	m := n - 1

	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		u[i]--
		v[i]--
		e[u[i]] = append(e[u[i]], edge{v[i], w[i]})
		e[v[i]] = append(e[v[i]], edge{u[i], w[i]})
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	var dfs func(idx, par int)
	dfs = func(idx, par int) {
		var costs []int
		for _, edge := range e[idx] {
			next := edge.to
			w := edge.w
			if next == par {
				continue
			}
			dfs(next, idx)
			costs = append(costs, dp[next][0]+w-dp[next][1])
			dp[idx][0] += dp[next][1]
			dp[idx][1] += dp[next][1]
		}
		sort.Slice(costs, func(i, j int) bool {
			return costs[i] > costs[j]
		})
		for i, cost := range costs {
			if cost <= 0 {
				break
			}
			if i < d[idx]-1 {
				dp[idx][0] += cost
			}
			if i < d[idx] {
				dp[idx][1] += cost
			}
		}
		if d[idx] == 0 {
			dp[idx][0] = -INF
		}
	}
	dfs(0, -1)
	//fmt.Println(dp)
	return dp[0][1]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	d := nextIntSlice(n)
	var u, v, w []int
	for i := 0; i < n-1; i++ {
		u = append(u, nextInt())
		v = append(v, nextInt())
		w = append(w, nextInt())
	}
	ans := solve(n, d, u, v, w)
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
