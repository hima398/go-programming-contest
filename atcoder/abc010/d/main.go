package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, g, e int, p, a, b []int) int {
	const INF = 1 << 60

	type edge struct {
		to, cap, rev int
	}
	edges := make([][]edge, n+1)
	for i := 0; i < e; i++ {
		edges[a[i]] = append(edges[a[i]], edge{b[i], 1, len(edges[b[i]])})
		edges[b[i]] = append(edges[b[i]], edge{a[i], 0, len(edges[a[i]]) - 1})

		edges[b[i]] = append(edges[b[i]], edge{a[i], 1, len(edges[a[i]])})
		edges[a[i]] = append(edges[a[i]], edge{b[i], 0, len(edges[b[i]]) - 1})
	}
	for _, v := range p {
		edges[v] = append(edges[v], edge{n, 1, len(edges[n])})
		edges[n] = append(edges[n], edge{v, 0, len(edges[v]) - 1})
	}

	var visited []bool
	var dfs func(i, t, f int) int
	dfs = func(i, t, f int) int {
		if i == t {
			return f
		}
		visited[i] = true
		for j, next := range edges[i] {
			if visited[next.to] {
				continue
			}
			if next.cap <= 0 {
				continue
			}
			d := dfs(next.to, t, next.cap)
			if d > 0 {
				edges[i][j].cap -= d
				edges[next.to][next.rev].cap += d
				return d
			}
		}
		return 0
	}
	ans := 0
	for {
		visited = make([]bool, n+1)
		f := dfs(0, n, INF)
		if f == 0 {
			return ans
		}
		ans += f
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, g, e := nextInt(), nextInt(), nextInt()
	p := nextIntSlice(g)
	var a, b []int
	for i := 0; i < e; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(n, g, e, p, a, b)
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
