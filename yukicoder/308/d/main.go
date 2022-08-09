package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a, b []int, c []string) string {
	uf := NewUnionFind(n)

	m := n * (n - 1) / 2
	for i := 0; i < m; i++ {
		uf.Unite(a[i], b[i])
		if uf.Size(1) == n {
			return c[i]
		}
	}
	return c[m-1]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var a, b []int
	var c []string
	for i := 0; i < n*(n-1)/2; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
		c = append(c, nextString())
	}
	ans := solve(n, a, b, c)
	PrintString(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
