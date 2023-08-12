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

	n, m := nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	k := nextInt()
	var x, y []int
	for i := 0; i < k; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt()-1)
	}
	lq := nextInt()
	var p, q []int
	for i := 0; i < lq; i++ {
		p = append(p, nextInt()-1)
		q = append(q, nextInt()-1)
	}
	ans := solve(n, m, u, v, k, x, y, lq, p, q)
	for _, v := range ans {
		Print(v)
	}
}

func solve(n, m int, u, v []int, k int, x, y []int, lq int, p, q []int) []string {
	//G初期状態の連結成分化する
	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		if uf.ExistSameUnion(u[i], v[i]) {
			continue
		}
		uf.Unite(u[i], v[i])
	}

	//良いグラフであることは、連結成分の代表点同士を結ばないようにする
	//→notConnect[pi][qi]がfalseになるようにデータ構造を準備する
	notConnect := make([]map[int]bool, n)
	for i := range notConnect {
		notConnect[i] = make(map[int]bool)
	}
	for i := 0; i < k; i++ {
		//xi, yiが含まれる連結成分の代表点、xRoot, yRoot
		xRoot, yRoot := uf.Find(x[i]), uf.Find(y[i])
		notConnect[xRoot][yRoot] = true
		notConnect[yRoot][xRoot] = true
	}

	var ans []string
	//各クエリの処理
	for i := 0; i < lq; i++ {
		//pi, qiが含まれる連結成分の代表点 pRoot, qRoot
		pRoot, qRoot := uf.Find(p[i]), uf.Find(q[i])
		if notConnect[pRoot][qRoot] {
			ans = append(ans, "No")
		} else {
			ans = append(ans, "Yes")
		}
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
