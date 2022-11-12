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

func solve(n, m int, x, y, a, b, z []int) int {
	const INF = 1 << 60

	for i := 0; i < m; i++ {
		a[i]--
		b[i]--
	}
	type road struct {
		s, t, cost int
	}

	f := func(n int, rs []road) int {
		var cost int
		uf := NewUnionFind(n + 2)
		//fmt.Println(len(rs), rs)
		for i := 0; i < len(rs); i++ {
			if !uf.ExistSameUnion(rs[i].s, rs[i].t) {
				cost += rs[i].cost
				uf.Unite(rs[i].s, rs[i].t)
			}
		}
		//PrintUnionFind(uf)
		if uf.Size(0) < n {
			return INF
		} else {
			return cost
		}
	}
	ans := INF
	for pat := 0; pat < 4; pat++ {
		var rs []road
		for i := 0; i < m; i++ {
			rs = append(rs, road{a[i], b[i], z[i]})
		}
		if pat&1 > 0 {
			for i := 0; i < n; i++ {
				rs = append(rs, road{i, n, x[i]})
			}
		}
		if pat&2 > 0 {
			for i := 0; i < n; i++ {
				rs = append(rs, road{i, n + 1, y[i]})
			}
		}

		sort.Slice(rs, func(i, j int) bool {
			return rs[i].cost < rs[j].cost
		})
		cost := f(n, rs)
		//fmt.Println(pat, cost)
		ans = Min(ans, cost)
		//ans = Min(ans, f(n+bits.OnesCount(uint(pat)), rs))
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	x := nextIntSlice(n)
	y := nextIntSlice(n)
	var a, b, z []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
		z = append(z, nextInt())
	}
	ans := solve(n, m, x, y, a, b, z)
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
