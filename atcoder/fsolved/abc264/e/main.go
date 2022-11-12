package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
}

func NewUnionFind(n, m int) *UnionFind {
	if n+m <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.parent = make([]int, n+m)
	u.rank = make([]int, n+m)
	u.size = make([]int, n+m)
	for i := 0; i < n+m; i++ {
		u.parent[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	for i := n; i < n+m; i++ {
		u.size[i] = 0
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

func solve(n, m, e int, u, v []int, q int, x []int) (ans []int) {
	//クエリで切っていく電線のID
	mx := make(map[int]struct{})
	for _, v := range x {
		mx[v] = struct{}{}
	}

	hasPowerPlant := make([]bool, n+m)
	for i := n; i < n+m; i++ {
		hasPowerPlant[i] = true
	}
	uf := NewUnionFind(n, m)

	for i := 0; i < e; i++ {
		//後からクエリで切られる電線はまだ処理しない
		if _, found := mx[i]; found {
			continue
		}
		//少なくてもどちらかの都市に電気が通っている
		b := hasPowerPlant[uf.Find(u[i])] || hasPowerPlant[uf.Find(v[i])]
		uf.Unite(u[i], v[i])
		if b {
			hasPowerPlant[uf.Find(u[i])] = true
		}
	}

	mr := make(map[int]int)
	for i := 0; i < n+m; i++ {
		mr[uf.Find(i)] = uf.Size(i)
	}
	//fmt.Println(mr)
	s := 0
	for k, v := range mr {
		if hasPowerPlant[k] {
			s += v
		}
	}

	var work []int
	work = append(ans, s)
	for i := q - 1; i >= 0; i-- {
		idx := x[i]
		if hasPowerPlant[uf.Find(u[idx])] && !hasPowerPlant[uf.Find(v[idx])] {
			s += uf.Size(v[idx])
			uf.Unite(u[idx], v[idx])
			hasPowerPlant[uf.Find(u[idx])] = true
		} else if hasPowerPlant[uf.Find(v[idx])] && !hasPowerPlant[uf.Find(u[idx])] {
			s += uf.Size(u[idx])
			uf.Unite(u[idx], v[idx])
			hasPowerPlant[uf.Find(v[idx])] = true
		} else {
			//どちらも都市のみ or どちらも電気が通っている
			uf.Unite(u[idx], v[idx])
		}
		work = append(work, s)
	}
	for i := len(work) - 2; i >= 0; i-- {
		ans = append(ans, work[i])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, e := nextInt(), nextInt(), nextInt()
	var u, v []int
	for i := 0; i < e; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	q := nextInt()
	var x []int
	for i := 0; i < q; i++ {
		x = append(x, nextInt()-1)
	}
	ans := solve(n, m, e, u, v, q, x)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
