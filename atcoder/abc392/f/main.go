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

	n := nextInt()
	p := nextIntSlice(n)

	ans := solve(n, p)

	PrintHorizonaly(ans)
}

func solve(n int, p []int) []int {
	tree := NewFenwickTree(n)
	for i := 0; i < n; i++ {
		tree.Update(i, 1)
	}

	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		idx := tree.IndexOf(p[i])
		ans[idx] = i + 1
		tree.Update(idx, -1)
	}

	return ans
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
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
	//fen.n = n + 1
	fen.n = 1
	for fen.n <= n+1 {
		fen.n <<= 1
	}
	fen.nodes = make([]int, fen.n)
	//bt.eval = f
	return fen
}

// i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	//内部では1-indexedなのでここでインクリメントする
	i++
	for i <= fen.n {
		fen.nodes[i] = fen.nodes[i] + v //fen.eval(fen.nodes[i], v)
		i += i & -i
	}
}

// i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	i++
	res := 0
	for i > 0 {
		res = fen.nodes[i] + res //fen.eval(fen.nodes[i], res)
		i -= i & -i
	}
	return res
}

func (fen *FenwickTree) IndexOf(k int) int {
	cur, w := 0, fen.n
	for w > 0 {
		if k > fen.nodes[cur+w] {
			cur, k = cur+w, k-fen.nodes[cur]
		}
		w >>= 1 // w /= 2
	}
	return cur
}
