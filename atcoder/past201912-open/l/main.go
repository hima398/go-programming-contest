package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var x, y, c []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
		c = append(c, nextInt())
	}
	var lx, ly, lc []int
	for i := 0; i < m; i++ {
		lx = append(lx, nextInt())
		ly = append(ly, nextInt())
		lc = append(lc, nextInt())
	}
	ans := solve(n, m, x, y, c, lx, ly, lc)
	PrintFloat64(ans)
}

type tower struct {
	idx     int
	x, y, c int
}

func computeDist2(x1, y1, x2, y2 int) float64 {
	xx := (x2 - x1)
	yy := (y2 - y1)
	return math.Sqrt(float64(xx*xx + yy*yy))
}
func computeCost(ts []tower) float64 {
	type node struct {
		u, v int
		c    float64
	}
	var costs []node
	for i := 0; i < len(ts)-1; i++ {
		for j := i + 1; j < len(ts); j++ {
			cost := computeDist2(ts[i].x, ts[i].y, ts[j].x, ts[j].y)
			if ts[i].c != ts[j].c {
				cost *= 10
			}
			costs = append(costs, node{ts[i].idx, ts[j].idx, cost})
		}
	}
	sort.Slice(costs, func(i, j int) bool {
		return costs[i].c < costs[j].c
	})
	uf := NewUnionFind(2 * len(ts))
	res := 0.0

	for len(costs) > 0 && uf.Size(0) < len(ts) {
		p := costs[0]
		costs = costs[1:]
		if uf.ExistSameUnion(p.u, p.v) {
			continue
		}
		res += p.c
		uf.Unite(p.u, p.v)
	}
	return res
}

func solve(n, m int, x, y, c, lx, ly, lc []int) float64 {
	var largeTowers, smallTowers []tower
	for i := 0; i < n; i++ {
		largeTowers = append(largeTowers, tower{i, x[i], y[i], c[i]})
	}
	for i := 0; i < m; i++ {
		smallTowers = append(smallTowers, tower{i + n, lx[i], ly[i], lc[i]})
	}
	ans := 1e18
	for pat := 0; pat < 1<<m; pat++ {
		ts := make([]tower, n)
		copy(ts, largeTowers)
		for i := 0; i < m; i++ {
			if pat>>i&1 == 1 {
				ts = append(ts, smallTowers[i])
			}
		}
		ans = math.Min(ans, computeCost(ts))
		//fmt.Println(cost)
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
	return i
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
