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

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, k := nextInt(), nextInt(), nextInt()
	var u, v, w []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		w = append(w, nextInt())
	}
	a := nextIntSlice(k)
	b := nextIntSlice(k)

	ans := solve(n, m, k, u, v, w, a, b)

	Print(ans)
}

func solve(n, m, k int, u, v, w, a, b []int) int {
	type edge struct {
		u, v, w int
	}
	var edges []edge
	for i := 0; i < m; i++ {
		edges = append(edges, edge{u[i], v[i], w[i]})
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	ma := make([]int, n)
	mb := make([]int, n)
	//最小全域木を作るためのUnionFind
	uf := NewUnionFind(n)
	//PrintUnionFind(uf)
	for i := 0; i < k; i++ {
		ma[a[i]-1]++
		mb[b[i]-1]++
	}

	//最小全域木の重みが小さい枝を追加しながら貪欲に解答を求める
	var ans int
	for _, e := range edges {
		if uf.ExistSameUnion(e.u, e.v) {
			continue
		}
		ru, rv := uf.Find(e.u), uf.Find(e.v)
		uf.Unite(e.u, e.v)
		root := uf.Find(e.u)
		ma[root] = ma[ru] + ma[rv]
		mb[root] = mb[ru] + mb[rv]

		x := Min(ma[root], mb[root])
		ans += e.w * x

		ma[root] -= x
		mb[root] -= x
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
