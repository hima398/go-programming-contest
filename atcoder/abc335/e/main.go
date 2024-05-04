package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/priorityqueue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := nextIntSlice(n)
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, m, a, u, v)
	Print(ans)
}

func solve(n, m int, a, u, v []int) int {
	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		if uf.ExistSameUnion(u[i], v[i]) {
			continue
		}
		if a[u[i]] == a[v[i]] {
			uf.Unite(u[i], v[i])
		}
	}
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		ud, vd := uf.Find(u[i]), uf.Find(v[i])
		if ud == vd {
			continue
		} else if a[ud] < a[vd] {
			e[ud] = append(e[ud], vd)
		} else {
			//a[ud]>a[vd]
			e[vd] = append(e[vd], ud)
		}
	}
	dp := make([]int, n)
	type node struct {
		i, v int
	}
	//q := queue.New[int]()
	q := priorityqueue.New[node](func(a, b node) int {
		if a.v == b.v {
			return 0
		}
		if a.v < b.v {
			return -1
		}
		return 1
	})

	q.Push(node{uf.Find(0), a[uf.Find(0)]})
	dp[uf.Find(0)] = 1
	for !q.Empty() {
		cur := q.Pop()
		for _, next := range e[cur.i] {
			if dp[next] > dp[cur.i] {
				continue
			}
			q.Push(node{next, a[next]})
			dp[next] = dp[cur.i] + 1
		}
	}
	//fmt.Println(dp)
	return dp[uf.Find(n-1)]
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
