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

	n, q := nextInt(), nextInt()
	var t []int
	u, v, k := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		switch t[i] {
		case 1:
			u[i], v[i] = nextInt()-1, nextInt()-1
		case 2:
			v[i], k[i] = nextInt()-1, nextInt()
		}
	}

	ans := solve(n, q, t, u, v, k)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, q int, t, u, v, k []int) []int {
	const MaxK = 10
	uf := NewUnionFind(n)
	ranking := make([][]int, n)
	for i := 0; i < n; i++ {
		ranking[i] = append(ranking[i], -i)
	}

	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			w, x := uf.Find(u[i]), uf.Find(v[i])
			if uf.ExistSameUnion(u[i], v[i]) {
				continue
			}
			uf.Unite(u[i], v[i])
			if uf.Find(u[i]) == w { // uが含まれていた集合にマージする
				ranking[w] = append(ranking[w], ranking[x]...)
				sort.Ints(ranking[w])
				if len(ranking[w]) > MaxK {
					ranking[w] = ranking[w][:MaxK]
				}
			} else { // uf.Find(u[i])==x, vが含まれていた集合にマージする
				ranking[x] = append(ranking[x], ranking[w]...)
				sort.Ints(ranking[x])
				if len(ranking[x]) > MaxK {
					ranking[x] = ranking[x][:MaxK]
				}
			}
		case 2:
			root := uf.Find(v[i])
			if len(ranking[root]) < k[i] {
				ans = append(ans, -1)
			} else {
				//fmt.Println(k[i], ranking[root])
				ans = append(ans, Abs(ranking[root][k[i]-1])+1)
			}
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
