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
	p := nextIntSlice(n - 1)
	for i := range p {
		p[i]--
	}
	q := nextInt()
	t, x, y := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		x[i] = nextInt() - 1
		if t[i] == 1 {
			y[i] = nextInt() - 1
		}
	}

	ans := solve(n, p, q, t, x, y)

	PrintVertically(ans)
}

func solve(n int, p []int, q int, t, x, y []int) []int {
	e := make([][]int, n)
	//クエリの際に辿りやすいように逆向きに結ぶ
	for i, pi := range p {
		e[i+1] = append(e[i+1], pi)
	}
	//fmt.Println(e)
	uf := NewUnionFind(n)
	var dfs func(cur, t int)
	dfs = func(cur, t int) {
		if cur == t {
			return
		}
		for _, nextPoint := range e[uf.Min(cur)] {
			nextUnit := uf.Min(nextPoint)
			uf.Unite(cur, nextUnit)
			//fmt.Println(cur, nextUnit, t)
			dfs(nextUnit, t)
		}
	}
	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			minX, minY := uf.Min(x[i]), uf.Min(y[i])
			dfs(minX, minY)
		case 2:
			ans = append(ans, uf.Min(x[i])+1)
		}
	}
	return ans
}

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	min    []int // min value in unit
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
	u.min = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.parent[i] = i
		u.rank[i] = 0
		u.min[i] = i
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

// xを含む連結成分のうち頂点の値が最小のものを返す
func (uf *UnionFind) Min(x int) int {
	return uf.min[uf.Find(x)]
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
		uf.min[y] = Min(uf.min[x], uf.min[y])
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		uf.min[x] = Min(uf.min[x], uf.min[y])
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
