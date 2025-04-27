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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a []int) int {
	var ans int

	graph := NewDirectedGraph(n)
	for i := 0; i < n; i++ {
		graph.AddEdge(i, a[i]-1)
		//自己ループしている点を先に数えておく
		if i == a[i]-1 {
			ans++
		}
	}

	for _, v := range graph.Scc().size {
		//サイズ2以上の強連結成分を答えに足す
		if v >= 2 {
			ans += v
		}
	}
	return ans
}

// 強連結成分分解用に
type DirectedGraph struct {
	n    int //頂点の数
	m    int //辺の数
	from []int
	to   []int
	size []int
}

func NewDirectedGraph(n int) *DirectedGraph {
	res := new(DirectedGraph)
	res.n = n
	res.size = make([]int, n)
	for i := range res.size {
		res.size[i] = 1
	}

	return res
}

func (g *DirectedGraph) AddEdge(from, to int) {
	g.from = append(g.from, from)
	g.to = append(g.to, to)
}

// 強連結成分分解して新しいグラフを返す
func (g *DirectedGraph) Scc() *DirectedGraph {
	label := make([]int, g.n)
	e := make([][]int, g.n)
	m := len(g.from)
	for i := 0; i < m; i++ {
		e[g.from[i]] = append(e[g.from[i]], g.to[i])
	}
	var idx int
	visited := make([]bool, g.n)
	for i := range label {
		label[i] = -1
	}
	var dfs func(cur int)
	dfs = func(cur int) {
		visited[cur] = true
		for _, next := range e[cur] {
			if visited[next] {
				continue
			}
			dfs(next)
		}
		label[idx] = cur
		idx++
	}
	for i := 0; i < g.n; i++ {
		if visited[i] {
			continue
		}
		dfs(i)
	}

	scc := make([]int, g.n)
	for i := range scc {
		scc[i] = -1
	}
	re := make([][]int, g.n)
	for i := 0; i < m; i++ {
		re[g.to[i]] = append(re[g.to[i]], g.from[i])
	}
	var revDfs func(cur, root int)
	revDfs = func(cur, root int) {
		scc[cur] = root
		for _, next := range re[cur] {
			if scc[next] >= 0 {
				continue
			}
			revDfs(next, root)
		}
	}

	for i := len(label) - 1; i >= 0; i-- {
		cur := label[i]
		if scc[cur] >= 0 {
			continue
		}
		revDfs(cur, cur)
	}

	res := NewDirectedGraph(g.n)
	for i := 0; i < m; i++ {
		if scc[g.from[i]] != scc[g.to[i]] {
			//異なる連結成分に辺を貼る
			res.AddEdge(g.from[i], g.to[i])
		}
	}
	for i, v := range scc {
		if i != v {
			res.size[v]++
			res.size[i]--
		}
	}
	return res
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
