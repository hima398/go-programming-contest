package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solveHonestly(n int, a []int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		ans = Max(ans, dp[i])
	}
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
		fen.nodes[i] = Max(fen.nodes[i], v)
		i += i & -i
	}
}

//i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	//i++
	res := 0
	for i > 0 {
		res = Max(fen.nodes[i], res)
		i -= i & -i
	}
	return res
}

func solve(n int, a []int) int {
	maxA := 0
	for _, ai := range a {
		maxA = Max(maxA, ai)
	}
	fTree := NewFenwickTree(maxA)
	for _, ai := range a {
		v := fTree.Query(ai - 1)
		fTree.Update(ai, v+1)
	}
	ans := fTree.Query(maxA)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	//ans := solveHonestly(n, a)
	ans := solve(n, a)
	PrintInt(ans)
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
