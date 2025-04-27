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
	a := nextIntSlice(n)

	//ans := solveHonestly(n, m, a)
	ans := solve(n, m, a)

	Print(ans)
}

func solve(n, m int, a []int) int {
	cumulativeSum := make([]int, n+1)
	for i := range a {
		cumulativeSum[i+1] = (cumulativeSum[i] + a[i]) % m
	}

	fenTree := NewFenwickTree(m)

	var s int
	var ans int
	for r := 1; r <= n; r++ {
		x := fenTree.Query(m) - fenTree.Query(cumulativeSum[r]+1)
		ans += cumulativeSum[r]*r - s + x*m
		s += cumulativeSum[r]
		fenTree.Update(cumulativeSum[r], 1)
	}

	return ans
}

func solveHonestly(n, m int, a []int) int {
	var ans int
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			var s int
			for k := l; k <= r; k++ {
				s += a[k]
			}
			ans += s % m
		}
	}
	return ans
}

type FenwickTree struct {
	n     int
	nodes []int
	//eval  func(x1, x2 int) int
}

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

// i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	//内部では1-indexedなのでここでインクリメントする
	i++
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
