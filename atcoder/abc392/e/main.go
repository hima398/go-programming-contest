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

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}

	ans := solve(n, m, a, b)

	Print(len(ans))
	if len(ans) > 0 {
		for _, v := range ans {
			PrintHorizonaly(v)
		}
	}
}

var need int

func solve(n, m int, a, b []int) [][3]int {
	need := n
	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		if uf.ExistSameUnion(a[i], b[i]) {
			//余ったケーブルとして記録しておく
			uf.Inc(i, a[i])
		} else {
			need--
			uf.Unite(a[i], b[i])
		}
	}

	type root struct {
		i, rem int
	}
	var rs []root
	for i := 0; i < n; i++ {
		if uf.Find(i) == i {
			rs = append(rs, root{i, uf.Rem(i)})
		}
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].rem > rs[j].rem
	})
	//fmt.Println(rs)
	r := rs[0].i
	var ans [][3]int
	for i := 1; i < len(rs); i++ {
		//ケーブルがないのに使おうとしている
		idx := uf.Dec(r)
		if idx < 0 {
			idx = uf.Dec(rs[i].i)
		}
		uf.Unite(r, rs[i].i)
		ans = append(ans, [3]int{idx + 1, b[idx] + 1, rs[i].i + 1})
		r = uf.Find(r)
	}

	return ans
}

type UnionFind struct {
	n      int
	parent []int // parentent numbers
	rank   []int // height of tree
	//rem    []int //余っているケーブル
	rem  [][]int
	size []int
}

func New(n int) *UnionFind {
	return NewUnionFind(n)
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	u.n = n
	// for accessing index without minus 1
	u.parent = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.rem = make([][]int, n+1)
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

func (uf *UnionFind) Rem(x int) int {
	return len(uf.rem[uf.Find(x)])
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
		//for i := 0; i < len(uf.rem[x]) && len(uf.rem[y]) < uf.n; i++ {
		//for i := 0; i < len(uf.rem[x]); i++ {
		uf.rem[y] = append(uf.rem[y], uf.rem[x]...)
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		//uf.rem[x] += uf.rem[y]
		//for i := 0; i < len(uf.rem[y]); i++ {
		//	for i := 0; i < len(uf.rem[y]) && len(uf.rem[x]) < uf.n; i++ {
		uf.rem[x] = append(uf.rem[x], uf.rem[y]...)

		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

// あまりのケーブルをインクリメントしておく
func (uf *UnionFind) Inc(i, x int) {
	idx := uf.Find(x)
	uf.rem[idx] = append(uf.rem[idx], i)
}

func (uf *UnionFind) Dec(x int) int {
	idx := uf.Find(x)
	if len(uf.rem[idx]) == 0 {
		return -1
	}

	ret := uf.rem[idx][0]
	uf.rem[idx] = uf.rem[idx][1:]
	return ret
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.parent)
	fmt.Println(u.rank)
	fmt.Println(u.size)
	fmt.Println(u.rem)
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

func PrintHorizonaly(x [3]int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
