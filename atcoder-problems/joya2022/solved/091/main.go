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

	n := nextInt()
	sx, sy, tx, ty := nextInt(), nextInt(), nextInt(), nextInt()
	var x, y, r []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
		r = append(r, nextInt())
	}
	ans := solve(n, sx, sy, tx, ty, x, y, r)
	PrintString(ans)
}

func computeDist2(x1, y1, x2, y2 int) int {
	xx := (x2 - x1)
	yy := (y2 - y1)
	return xx*xx + yy*yy
}

func solve(n, sx, sy, tx, ty int, x, y, r []int) string {
	uf := NewUnionFind(n)
	check := func(x1, y1, r1, x2, y2, r2 int) bool {
		dist := computeDist2(x1, y1, x2, y2)
		//外部で接する
		ok := (r1+r2)*(r1+r2) == dist
		//交わる
		if r1 <= r2 {
			//大きい方の円に含まれる
			ok = ok || (r2-r1)*(r2-r1) <= dist && dist < (r2+r1)*(r2+r1)
		} else {
			ok = ok || (r1-r2)*(r1-r2) <= dist && dist < (r1+r2)*(r1+r2)
		}
		return ok
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if check(x[i], y[i], r[i], x[j], y[j], r[j]) {
				uf.Unite(i, j)
			}
		}
	}
	s, t := -1, -1
	for i := 0; i < n; i++ {
		dist := computeDist2(sx, sy, x[i], y[i])
		if dist == r[i]*r[i] {
			s = i
		}
	}
	for i := 0; i < n; i++ {
		dist := computeDist2(tx, ty, x[i], y[i])
		if dist == r[i]*r[i] {
			t = i
		}
	}
	if uf.ExistSameUnion(s, t) {
		return "Yes"
	} else {
		return "No"
	}
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
