package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type FenwickTree struct {
	n     int
	nodes []int
	//eval  func(x1, x2 int) int
}

func NewFenwickTree(n int) *FenwickTree {
	fen := new(FenwickTree)
	// 1-indexed
	fen.n = n + 1
	fen.nodes = make([]int, fen.n)
	//bt.eval = f
	return fen
}

// i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	//内部では1-indexedなのでここでインクリメントする
	//i++
	for i < fen.n {
		fen.nodes[i] = fen.nodes[i] + v //fen.eval(fen.nodes[i], v)
		i += i & -i
	}
}

// i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	//i++
	res := 0
	for i > 0 {
		res = fen.nodes[i] + res //fen.eval(fen.nodes[i], res)
		i -= i & -i
	}
	return res
}

func (fen *FenwickTree) RangeSum(l, r int) int {
	return fen.Query(r) - fen.Query(l-1)
}

func solveWithFenwick(n, q int, c, l, r []int) []int {
	for i := 0; i < n; i++ {
		//c[i]--
	}
	type Query struct {
		i, l, r int
	}
	//q1, q2 := make([][]int, n), make([][]int, n)
	qs := make([][]Query, n+1)

	for i := 0; i < q; i++ {
		//l[i]--
		//r[i]--
		//q1[r[i]] = append(q1[r[i]], l[i])
		//q2[r[i]] = append(q2[r[i]], i)
		qs[r[i]] = append(qs[r[i]], Query{i, l[i], r[i]})
	}
	a := make([]int, n+1)
	ans := make([]int, q)
	fenwick := NewFenwickTree(n)
	for i := 1; i <= n; i++ {
		if a[c[i-1]] > 0 {
			fenwick.Update(a[c[i-1]], -1)
		}
		fenwick.Update(i, 1)
		a[c[i-1]] = i
		for j := range qs[i] {
			ans[qs[i][j].i] = fenwick.RangeSum(qs[i][j].l, i)
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	c := nextIntSlice(n)
	var l, r []int
	for i := 0; i < q; i++ {
		l = append(l, nextInt())
		r = append(r, nextInt())
	}
	//ans := solveWithFenwick(n, q, c, l, r)
	ans := solveWithMo(n, q, c, l, r)
	PrintVertically(ans)
}

// solve with Mo's algorithm
func solveWithMo(n, q int, c, l, r []int) []int {
	d := n/Min(n, Sqrt(q+1)) + 1
	type query struct {
		i, l, r int
	}
	var qs []query
	for i := 0; i < q; i++ {
		qs = append(qs, query{i, l[i] - 1, r[i]})
	}
	sort.Slice(qs, func(i, j int) bool {
		li, lj := qs[i].l/d, qs[j].l/d
		if li == lj {
			if li%2 == 0 {
				return qs[i].r < qs[j].r
			} else {
				return qs[i].r > qs[j].r
			}
		}
		return li < lj
	})
	//fmt.Println(qs)
	colors := make([]int, n+1) //make(map[int]int)
	var cur int
	add := func(c int) {
		if colors[c] == 0 {
			cur++
		}
		colors[c]++
	}
	del := func(c int) {
		colors[c]--
		if colors[c] == 0 {
			cur--
		}
	}
	var curL, curR int
	ans := make([]int, q)
	for _, v := range qs {
		nl, nr := v.l, v.r
		for curR < nr {
			add(c[curR])
			curR++
		}
		for nl < curL {
			curL--
			add(c[curL])
		}
		for nr < curR {
			curR--
			del(c[curR])
		}
		for curL < nl {
			del(c[curL])
			curL++
		}
		ans[v.i] = cur
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
	if x > y {
		return y
	}
	return x
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}
