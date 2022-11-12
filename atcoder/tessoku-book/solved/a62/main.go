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

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(n, m, a, b)
	PrintString(ans)
}

func solve(n, m int, a, b []int) string {
	g := NewGraph(n, m)
	for i := 0; i < g.m; i++ {
		g.Add(a[i], b[i])
	}
	g.Dfs(0)
	ok := true
	for i := 0; i < n; i++ {
		ok = ok && g.v[i]
	}
	if ok {
		return "The graph is connected."
	} else {
		return "The graph is not connected."
	}
}

type Graph struct {
	n, m int
	v    []bool  // visited
	e    [][]int //edges
}

func New(n, m int) *Graph {
	return NewGraph(n, m)
}

func NewGraph(n, m int) *Graph {
	g := new(Graph)
	g.n = n
	g.m = m
	g.v = make([]bool, n)
	g.e = make([][]int, n)
	return g
}

func (g *Graph) Add(a, b int) {
	g.e[a] = append(g.e[a], b)
	g.e[b] = append(g.e[b], a)
}

func (g *Graph) Dfs(cur int) {
	g.v[cur] = true
	for _, next := range g.e[cur] {
		if g.v[next] {
			continue
		}
		g.Dfs(next)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
