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
	for i := range a {
		a[i]--
	}

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a []int) int {
	uf1 := NewUnionFind(n)
	uf2 := NewUnionFind(n)
	foundCycle := make([]bool, n)
	for i := 0; i < n; i++ {
		cnt := make(map[int]int)
		cur := i
		for {
			if foundCycle[uf1.Find(cur)] {
				break
			}
			cnt[cur]++
			if cnt[cur] >= 3 {
				foundCycle[uf1.Find(cur)] = true
				break
			}
			next := a[cur]
			uf1.Unite(cur, next)
			if cnt[next] == 1 {
				if !uf2.ExistSameUnion(cur, next) {
					uf2.Unite(cur, next)
				}
			}
			cur = next
		}
		//fmt.Println(visited)
		//fmt.Println(cnt)
	}

	e := make([][]int, n)
	for i, ai := range a {
		if uf2.ExistSameUnion(i, ai) {
			continue
		}
		from, to := uf2.Find(i), uf2.Find(ai)
		e[from] = append(e[from], to)
	}
	memo := make([]int, n)
	var dfs func(cur int) int
	dfs = func(cur int) int {
		if memo[cur] > 0 {
			return memo[cur]
		}
		var s int
		for _, next := range e[cur] {
			s += dfs(next)
		}
		memo[cur] = s + uf2.Size(cur)
		return memo[cur]
	}
	for i := 0; i < n; i++ {
		if uf2.Find(i) != i {
			continue
		}
		dfs(i)
	}
	var ans int
	for i, v := range memo {
		ans += uf2.Size(i) * v
	}
	return ans
}

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
}

func New(n int) *UnionFind {
	return NewUnionFind(n)
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.parent = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.parent[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	return u
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] == x {
		return x
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4
		uf.parent[x] = uf.Find(uf.parent[x])
		return uf.parent[x]
	}
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Find(x)]
}

func (uf *UnionFind) ExistSameUnion(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	// rank
	if uf.rank[x] < uf.rank[y] {
		//yがrootの木にxがrootの木を結合する
		uf.parent[x] = y
		uf.size[y] += uf.size[x]
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.parent)
	fmt.Println(u.rank)
	fmt.Println(u.size)
}

func solveScc(n int, a []int) int {
	g1 := NewDirectedGraph(n)
	for from, to := range a {
		g1.AddEdge(from, to)
	}

	g2 := g1.Scc()

	d := make([]int, n)
	for i := 0; i < g2.m; i++ {
		d[g2.to[i]]++
	}
	visited := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {

	}
	var ans int
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		dfs(i)
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
