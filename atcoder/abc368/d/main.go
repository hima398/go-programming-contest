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

	n, k := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	v := nextIntSlice(k)
	for i := range v {
		v[i]--
	}

	ans := solve(n, k, a, b, v)

	Print(ans)
}

func solve(n, k int, a, b, v []int) int {
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}

	vm := make(map[int]struct{})
	for _, vi := range v {
		vm[vi] = struct{}{}
	}

	uf := NewUnionFind(n)
	for _, vi := range v {
		uf.vsize[vi]++
	}

	var dfs func(cur, par int) bool
	dfs = func(cur, par int) bool {
		var connected bool
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			connected = dfs(next, cur) || connected
		}
		var ok bool
		if _, found := vm[cur]; found && par >= 0 && uf.Vsize(cur) < len(v) {
			uf.Unite(par, cur)
			ok = true
		} else if connected && par >= 0 && uf.Vsize(cur) < len(v) {
			uf.Unite(par, cur)
			ok = true
		}
		return ok
	}

	dfs(0, -1)
	ans := uf.Size(v[0])

	return ans
}

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
	vsize  []int
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
	u.vsize = make([]int, n+1)
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

func (uf *UnionFind) Vsize(x int) int {
	return uf.vsize[uf.Find(x)]
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
		uf.vsize[y] += uf.vsize[x]
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		uf.vsize[x] += uf.vsize[y]
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
	fmt.Println(u.vsize)
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
