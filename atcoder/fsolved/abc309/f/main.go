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
	var h, w, d []int
	for i := 0; i < n; i++ {
		h = append(h, nextInt())
		w = append(w, nextInt())
		d = append(d, nextInt())
	}
	//ok := solve(n, h, w, d)
	ok := solve01(n, h, w, d)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(n int, h, w, d []int) bool {
	type box struct {
		m []int
	}
	var bs []box
	for i := 0; i < n; i++ {
		bs = append(bs, box{[]int{h[i], w[i], d[i]}})
		sort.Ints(bs[i].m)
	}
	sort.Slice(bs, func(i, j int) bool {
		if bs[i].m[0] == bs[j].m[0] {
			return bs[i].m[1] > bs[j].m[1]
		}
		return bs[i].m[0] < bs[j].m[0]
	})
	var ws []int
	for _, bx := range bs {
		ws = append(ws, bx.m[1])
	}
	compress := NewCompress()
	compress.Init(ws)
	segTree := NewSegmentTree(n, 1<<60, func(x1 int, x2 int) int {
		if x1 < x2 {
			return x1
		}
		return x2
	})
	for _, bx := range bs {
		idx := compress.GetIndex(bx.m[1])
		if segTree.Query(0, idx) < bx.m[2] {
			return true
		}
		segTree.Update(idx, bx.m[2])
	}
	return false
}

func solve01(n int, h, w, d []int) bool {
	type box struct {
		m []int
	}
	var bs []box
	for i := 0; i < n; i++ {
		bs = append(bs, box{[]int{h[i], w[i], d[i]}})
		sort.Ints(bs[i].m)
	}
	sort.Slice(bs, func(i, j int) bool {
		if bs[i].m[0] == bs[j].m[0] {
			return bs[i].m[1] > bs[j].m[1]
		}
		return bs[i].m[0] < bs[j].m[0]
	})
	var ws []int
	for _, bx := range bs {
		ws = append(ws, bx.m[1])
	}
	compress := NewCompress()
	compress.Init(ws)
	fenwick := NewFenwickTree(n)
	for _, bx := range bs {
		idx := compress.GetIndex(bx.m[1])
		if fenwick.Query(idx) < bx.m[2] {
			return true
		}
		fenwick.Update(idx, bx.m[2])
	}
	return false
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
	for i := range fen.nodes {
		fen.nodes[i] = 1 << 60
	}
	//bt.eval = f
	return fen
}

// i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	//内部では1-indexedなのでここでインクリメントする
	i++
	for i < fen.n {
		fen.nodes[i] = Min(fen.nodes[i], v) //fen.nodes[i] + v //fen.eval(fen.nodes[i], v)
		i += i & -i
	}
}

// i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	i++
	res := 1 << 60
	for i > 0 {
		res = Min(res, fen.nodes[i]) //fen.nodes[i] + res //fen.eval(fen.nodes[i], res)
		i -= i & -i
	}
	return res
}

type SegmentTree struct {
	size  int
	nodes []int
	f     func(x1, x2 int) int
	inf   int
}

func NewSegmentTree(n, inf int, f func(x1, x2 int) int) (st *SegmentTree) {
	st = new(SegmentTree)
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([]int, 2*st.size)
	for i := range st.nodes {
		st.nodes[i] = inf
	}
	st.inf = inf
	st.f = f
	return st
}

func (this *SegmentTree) QueryRecursively(a, b, k, l, r int) int {
	// [a, b)と[l, r)が交差しない
	if a >= r || b <= l {
		return this.inf
	}

	// [a, b)が[l, r)を完全に含んでいる
	if a <= l && b >= r {
		return this.nodes[k]
	}

	vl := this.QueryRecursively(a, b, 2*k, l, (l+r)/2)
	vr := this.QueryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return this.f(vl, vr) //Max(vl, vr)
}

// [l, r)の区間の値をxorした結果を返す
func (this *SegmentTree) Query(l, r int) int {
	return this.QueryRecursively(l, r, 1, 0, this.size)
}

func (this *SegmentTree) Update(k, x int) {
	k += this.size
	this.nodes[k] = this.f(this.nodes[k], x)
	for k > 1 {
		k /= 2
		this.nodes[k] = this.f(this.nodes[k*2], this.nodes[k*2+1])
	}
	//fmt.Println(this.nodes)
}

type Compress struct {
	//重複除去済みの圧縮元
	x []int
}

func New() *Compress {
	return NewCompress()
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

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
