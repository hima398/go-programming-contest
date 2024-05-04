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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintVertically(ans)
}

func solve(n int, a []int) []int {
	const p = 998244353

	wTree := NewFenwickTree(2*int(1e5)+1, p)
	sTree := NewFenwickTree(2*int(1e5)+1, p)
	var ans []int
	var v int
	for i := 0; i < n; i++ {
		v += 2 * (wTree.Query(a[i]) * a[i]) % p
		v %= p
		v += 2 * (sTree.Query(2*int(1e5)) - sTree.Query(a[i]) + p) % p
		v %= p
		v += a[i]
		v %= p

		in := Inv(i+1, p)
		in *= in
		in %= p

		ans = append(ans, (v*in)%p)

		wTree.Update(a[i], 1)
		sTree.Update(a[i], a[i])
	}
	return ans
}

type FenwickTree struct {
	n, p  int
	nodes []int
	//eval  func(x1, x2 int) int
}

func NewFenwickTree(n, p int) *FenwickTree {
	fen := new(FenwickTree)
	// 1-indexed
	fen.n = n + 1
	fen.p = p
	fen.nodes = make([]int, fen.n)
	//bt.eval = f
	return fen
}

//i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	//内部では1-indexedなのでここでインクリメントする
	//i++
	for i < fen.n {
		fen.nodes[i] = (fen.nodes[i] + v) % fen.p //fen.eval(fen.nodes[i], v)
		i += i & -i
	}
}

//i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	//i++
	res := 0
	for i > 0 {
		res = (fen.nodes[i] + res) % fen.p //fen.eval(fen.nodes[i], res)
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
