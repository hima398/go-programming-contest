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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(n)
	ans := solve(n, a, b)
	PrintInt(ans)
}

func computeLis(n int, p []int) int {
	fen := NewFenwickTree(n)
	for i := 0; i < n; i++ {
		max := fen.Query(p[i] - 1)
		fen.Update(p[i], max+1)
	}
	//fmt.Println(fen)
	return fen.Query(n)
}

func solve(n int, a, b []int) int {
	type perm struct {
		a, b int
	}
	var ps []perm
	for i := 0; i < n; i++ {
		ps = append(ps, perm{a[i], b[i]})
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].a < ps[j].a
	})
	for i := 0; i < n; i++ {
		b[i] = ps[i].b
	}
	ans := computeLis(n, b)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].b < ps[j].b
	})
	for i := 0; i < n; i++ {
		a[i] = ps[i].a
	}
	ans = Max(ans, computeLis(n, a))
	ans += n

	return ans
}

type FenwickTree struct {
	n     int
	nodes []int
	//eval  func(x1, x2 int) int
}

//
func New(n int) *FenwickTree {
	return NewFenwickTree(n)
}

func NewFenwickTree(n int) *FenwickTree {
	fen := new(FenwickTree)
	// 1-indexed
	fen.n = n + 1
	fen.nodes = make([]int, fen.n)
	//bt.eval = f
	return fen
}

//i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	//内部では1-indexedなのでここでインクリメントする
	//i++
	for i < fen.n {
		fen.nodes[i] = Max(v, fen.nodes[i])
		i += i & -i
	}
}

//i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	//i++
	res := 0
	for i > 0 {
		res = Max(res, fen.nodes[i])
		i -= i & -i
	}
	return res
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
