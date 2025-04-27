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
	var u, v []int
	for i := 0; i < n-1; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}

	ans := solve(n, u, v)

	Print(ans)
}

func solve(n int, u, v []int) int {
	d := make([]int, n)
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		d[u[i]]++
		d[v[i]]++
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	//fmt.Println(d)

	uf := NewUnionFind(n)

	var dfs func(cur, par int)
	dfs = func(cur, par int) {
		if par >= 0 && d[cur] == 3 && d[par] == 3 {
			uf.Unite(cur, par)
		}
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			dfs(next, cur)
		}
	}
	dfs(0, -1)

	//次数3の頂点のグループ近傍にある次数2の頂点の個数
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		if d[i] != 3 {
			continue
		}
		for _, near := range e[i] {
			if d[near] == 2 {
				cnt[uf.Find(i)]++
			}
		}
	}
	//fmt.Println(cnt)

	var ans int
	for _, v := range cnt {
		ans += v * (v - 1) / 2
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
