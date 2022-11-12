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
	PrintVertically(ans)
}

func solve(n, m int, a, b []int) []int {
	g := NewGraph(n, m)
	for i := 0; i < m; i++ {
		g.Add(a[i], b[i])
	}
	g.Bfs(0)

	return g.d
}

type Graph struct {
	n, m int
	v    []bool // visited
	d    []int
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
	g.d = make([]int, n)
	for i := 0; i < n; i++ {
		g.d[i] = -1
	}
	g.e = make([][]int, n)
	return g
}

func (g *Graph) Add(a, b int) {
	g.e[a] = append(g.e[a], b)
	g.e[b] = append(g.e[b], a)
}

func (g *Graph) Bfs(s int) {
	var q []int
	q = append(q, s)
	g.v[s] = true
	g.d[s] = 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range g.e[cur] {
			if g.v[next] {
				continue
			}
			q = append(q, next)
			g.v[next] = true
			g.d[next] = g.d[cur] + 1
		}
	}
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
