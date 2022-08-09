package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a, b []int) ([]int, error) {
	type edge struct {
		i, t int
	}
	e := make([][]edge, n+1)
	uf := NewUnionFind(n)
	for i := range a {
		//e = append(e, edge{i, a[i], b[i]})
		e[a[i]] = append(e[a[i]], edge{i, b[i]})
		e[b[i]] = append(e[b[i]], edge{i, a[i]})
		uf.Unite(a[i], b[i])
	}
	visited := make([]bool, n+1)
	visited[0] = true
	ans := make([]int, n)
	var dfs func(cur, par int) int
	dfs = func(cur, par int) int {
		if visited[cur] {
			return 0
		}
		visited[cur] = true
		res := 1
		for _, next := range e[cur] {
			if next.t == par {
				continue
			}
			if ans[next.i] == 0 {
				ans[next.i] = next.t
			}

			res += dfs(next.t, cur)
		}
		return res
	}
	ok := true
	for i := 1; i <= n; i++ {
		if visited[uf.Find(i)] {
			continue
		}
		ne := dfs(uf.Find(i), 0)
		fmt.Println(uf.Find(i), uf.Size(i), ne)
		ok = ok && uf.Size(i) == ne
	}
	for _, b := range visited {
		ok = ok && b
	}
	m := make(map[int]struct{})
	for _, v := range ans {
		m[v] = struct{}{}
	}
	if ok && len(m) == n {
		return ans, nil
	} else {
		return nil, errors.New("No")
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans, err := solve(n, a, b)
	if err != nil {
		PrintString("No")
		return
	}
	PrintString("Yes")
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

type UnionFind struct {
	par  []int // parent numbers
	rank []int // height of tree
	size []int
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.par = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.par[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	return u
}

func (this *UnionFind) Find(x int) int {
	if this.par[x] == x {
		return x
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4
		this.par[x] = this.Find(this.par[x])
		return this.par[x]
	}
}

func (this *UnionFind) Size(x int) int {
	return this.size[this.Find(x)]
}

func (this *UnionFind) ExistSameUnion(x, y int) bool {
	return this.Find(x) == this.Find(y)
}

func (this *UnionFind) Unite(x, y int) {
	x = this.Find(x)
	y = this.Find(y)
	if x == y {
		return
	}
	// rank
	if this.rank[x] < this.rank[y] {
		//yがrootの木にxがrootの木を結合する
		this.par[x] = y
		this.size[y] += this.size[x]
	} else {
		// this.rank[x] >= this.rank[y]
		//xがrootの木にyがrootの木を結合する
		this.par[y] = x
		this.size[x] += this.size[y]
		if this.rank[x] == this.rank[y] {
			this.rank[x]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.par)
	fmt.Println(u.rank)
	fmt.Println(u.size)
}
