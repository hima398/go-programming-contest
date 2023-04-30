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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(n)

	ans := solve(n, a, b)

	PrintInt(ans)
}

func solve(n int, a, b []int) int {
	m := make(map[int]int)
	ca := NewCompress()
	ca.Init(a)
	d := make([][]int, ca.Size())
	for i := range b {
		b[i] = -b[i]
	}
	for i := 0; i < n; i++ {
		idx := ca.GetIndex(a[i])
		d[idx] = append(d[idx], b[i])
		m[b[i]] = 0
	}
	cb := NewCompress()
	cb.Init(b)
	for k := range m {
		m[k] = cb.GetIndex(k)
	}
	//fmt.Println(m)
	//fmt.Println(d)

	fen := NewFenwickTree(len(m))
	var ans int
	for _, bs := range d {
		for _, v := range bs {
			fen.Update(m[v], 1)
		}
		for _, v := range bs {
			ans += fen.Query(m[v] + 1)
		}
	}
	//fmt.Println(fen)
	return ans
}

type Compress struct {
	//重複除去済みの圧縮元
	x []int
}

func NewCompress() *Compress {
	return new(Compress)
}

func (c *Compress) Init(x []int) {
	m := make(map[int]struct{})
	for _, v := range x {
		m[v] = struct{}{}
	}
	for k := range m {
		c.x = append(c.x, k)
	}
	sort.Ints(c.x)
}

func (c *Compress) GetIndex(x int) int {
	return sort.Search(len(c.x), func(i int) bool {
		return c.x[i] >= x
	})
}

func (c *Compress) Size() int {
	return len(c.x)
}

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
	i++
	for i < fen.n {
		fen.nodes[i-1] += v //fen.eval(fen.nodes[i], v)
		i += i & -i
	}
}

// i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	//i++
	res := 0
	for i > 0 {
		res += fen.nodes[i-1] //fen.eval(fen.nodes[i], res)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
