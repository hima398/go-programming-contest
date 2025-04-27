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
	var a, b, d []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		d = append(d, nextInt())
	}

	ans := solve(n, q, a, b, d)

	if len(ans) > 0 {
		PrintHorizonaly(ans)
	}
}

func solve(n, q int, a, b, d []int) []int {
	const INF = 1 << 60

	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	var ans []int
	uf := NewUnionFind(n)
	for i := 0; i < q; i++ {
		if uf.ExistSameUnion(a[i], b[i]) {
			//fmt.Printf("w(a) = %d, w(b) = %d, diff = %d, d = %d\n", uf.Weight(a[i]), uf.Weight(b[i]), uf.Diff(a[i], b[i]), d[i])
			//PrintUnionFind(uf)
			if uf.Diff(a[i], b[i]) == d[i] {
				ans = append(ans, i+1)
			}
		} else {
			uf.Unite(a[i], b[i], d[i])
			ans = append(ans, i+1)
		}
		//PrintUnionFind(uf)
	}

	return ans
}

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
	weight []int // 親ノードとの値の差分
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
	u.weight = make([]int, n+1)
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
		parent := uf.Find(uf.parent[x])
		uf.weight[x] += uf.weight[uf.parent[x]]
		uf.parent[x] = parent
		return uf.parent[x]
	}
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Find(x)]
}

func (uf *UnionFind) Weight(x int) int {
	uf.Find(x)
	return uf.weight[x]
}

func (uf *UnionFind) Diff(x, y int) int {
	return uf.Weight(y) - uf.Weight(x)
}

func (uf *UnionFind) ExistSameUnion(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Unite(x, y, w int) {
	w += uf.Weight(x)
	w -= uf.Weight(y)

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
		uf.weight[x] = -w
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		uf.weight[y] = w
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
	fmt.Println(u.weight)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
