package main

import (
	"bufio"
	"errors"
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
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, m, u, v)
	//ans := solveCommentary(n, m, u, v)
	PrintInt(ans)
}

func solveCommentary(n, m int, u, v []int) int {
	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		if uf.ExistSameUnion(u[i], v[i]) {
			continue
		}
		uf.Unite(u[i], v[i])
	}
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	g := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = -1
	}

	v1 := make([]bool, n)

	bfs := func(root int) (b, w int, err error) {
		var q []int
		q = append(q, root)
		v1[root] = true
		g[root] = 0
		b++
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				if v1[next] {
					if g[next] == g[cur] {
						return -1, -1, errors.New("Impossible")
					}
					continue
				}
				q = append(q, next)
				v1[next] = true
				g[next] = g[cur] ^ 1
				if g[next] == 0 {
					b++
				} else if g[next] == 1 {
					w++
				}
			}
		}
		return b, w, nil
	}
	ans := n*(n-1)/2 - m
	for i := 0; i < n; i++ {
		if i != uf.Find(i) {
			continue
		}
		b, w, err := bfs(i)
		if err != nil {
			//ans -= uf.Size(i) * (n - 1) / 2
			return 0
		} else {
			ans -= b * (b - 1) / 2
			ans -= w * (w - 1) / 2
		}
	}
	return ans
}

func solve(n, m int, u, v []int) int {
	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		if uf.ExistSameUnion(u[i], v[i]) {
			continue
		}

		uf.Unite(u[i], v[i])
	}
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = len(e[i])
	}
	g := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = -1
	}
	v1 := make([]bool, n)
	bfs := func(root int) (bool, []int) {
		res := make([]int, 2)
		var q []int
		q = append(q, root)
		g[root] = 0
		v1[root] = true
		//res[0]++
		ok := true
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			res[g[cur]]++
			for _, next := range e[cur] {
				if v1[next] {
					if g[next] == g[cur] {
						//二部グラフじゃない
						//return false, nil
						ok = false
						continue
					} else {
						continue
					}
				}
				q = append(q, next)
				g[next] = g[cur] ^ 1
				v1[next] = true
				//res[g[next]]++
			}
		}
		if ok {
			return ok, res
		} else {
			return ok, nil
		}
	}
	s := n
	type graph struct {
		root  int
		group []int
	}
	var gg []graph
	for i := 0; i < n; i++ {
		if uf.Find(i) != i {
			continue
		}
		isBipartite, group := bfs(i)
		if !isBipartite {
			//n -= uf.Size(i)
			return 0
		} else {
			gg = append(gg, graph{i, group})
		}
	}
	//fmt.Println(gg)
	var ans int
	visited := make([]bool, n)
	bfs2 := func(root int, group []int) int {
		res := 0
		var q []int
		q = append(q, root)
		visited[root] = true
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			res += group[g[cur]^1] - d[cur]
			//fmt.Println(cur, res)
			for _, next := range e[cur] {
				if visited[next] {
					continue
				}
				q = append(q, next)
				visited[next] = true
			}
		}
		return res
	}
	for _, v := range gg {
		ans += uf.Size(v.root) * (s - uf.Size(v.root))
		ans += bfs2(v.root, v.group)
	}
	return ans / 2
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
