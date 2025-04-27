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

	n, q := nextInt(), nextInt()
	t, x, c := make([]int, q), make([]int, q), make([]int, q)

	for i := 0; i < q; i++ {
		t[i] = nextInt()
		if t[i] == 1 {
			x[i] = nextInt() - 1
		}
		c[i] = nextInt() - 1
	}
	ans := solve(n, q, t, x, c)

	PrintVertically(ans)
}

func solve(n, q int, t, x, c []int) []int {
	//色cの個数
	color := make([]int, n)
	for i := range color {
		color[i] = 1
	}
	//連結成分の左端、右端
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		l[i], r[i] = i, i
	}
	uf := NewUnionFind(n)

	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			s := uf.Size(x[i])
			color[uf.Color(x[i])] -= s
			//色を塗り替える
			uf.Paint(x[i], c[i])
			color[c[i]] += s

			//左側と同じ色になれば連結する
			prev := uf.Left(x[i]) - 1
			if prev >= 0 && uf.Color(prev) == uf.Color(x[i]) {
				uf.Unite(prev, x[i])
			}
			//右側と同じ色になれば連結する
			next := uf.Right(x[i]) + 1
			if next < n && uf.Color(x[i]) == uf.Color(next) {
				uf.Unite(x[i], next)
			}
		case 2:
			ans = append(ans, color[c[i]])
		}
	}
	return ans
}

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
	//extend
	l, r  []int
	color []int
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
	u.l, u.r = make([]int, n+1), make([]int, n+1)
	u.color = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.parent[i] = i
		u.rank[i] = 0
		u.size[i] = 1
		u.l[i], u.r[i] = i, i
		u.color[i] = i
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

func (uf *UnionFind) Left(x int) int {
	return uf.l[uf.Find(x)]
}
func (uf *UnionFind) Right(x int) int {
	return uf.r[uf.Find(x)]
}

func (uf *UnionFind) Color(x int) int {
	return uf.color[uf.Find(x)]
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
		uf.l[y] = Min(uf.l[x], uf.l[y])
		uf.r[y] = Max(uf.r[x], uf.r[y])
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		uf.l[x] = Min(uf.l[x], uf.l[y])
		uf.r[x] = Max(uf.r[x], uf.r[y])
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

func (uf *UnionFind) Paint(x, c int) {
	uf.color[uf.Find(x)] = c
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
