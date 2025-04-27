package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
	a := nextIntSlice(n)

	ans := solveHonestly(n, a)

	Print(ans)
}

func solveHonestly(n int, a []int) int {
	l, r := make([]int, n), make([]int, n)
	ml := make(map[int]int)
	for i := range a {
		ml[a[i]]++
		l[i] = len(ml)
	}
	mr := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		mr[a[i]]++
		r[i] = len(mr)
	}
	//fmt.Println(l, r)

	x := make([]int, n)
	last := make(map[int]int)
	dp := NewLazySegmentTree(n)
	e := -int(1e9)
	op := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	mapping := func(f, x int) int {
		return f + x
	}
	composition := func(f, g int) int {
		return f + g
	}
	dp.Init(e, 0, op, mapping, composition)
	for i := 0; i < n; i++ {
		var j int
		if v, found := last[a[i]]; found {
			j = v
		}
		dp.Apply(j, i, 1)
		x[i] = dp.Prod(0, i)
		dp.Set(i, l[i])
		last[a[i]] = i
		//fmt.Println("i = ", i)
		//fmt.Println(dp.node)
		//fmt.Println(dp.lazy)
		//fmt.Println(x)
	}
	var ans int
	for i := 1; i < n-1; i++ {
		ans = Max(ans, x[i]+r[i+1])
	}
	return ans
}

type LazySegmentTree struct {
	n           int
	size        int
	log         int
	e           int //二項演算における単位元
	inf         int
	node        []int
	lazy        []int
	op          func(x, y int) int
	mapping     func(x, y int) int
	composition func(x, y int) int
}

func NewLazySegmentTree(n int) *LazySegmentTree {
	res := new(LazySegmentTree)
	res.n = n
	res.size = 1
	for res.size < res.n {
		res.size *= 2
	}
	res.node = make([]int, 2*res.size)
	res.lazy = make([]int, 2*res.size)

	res.log = bits.TrailingZeros(uint(res.size))

	return res
}

func (tree *LazySegmentTree) Init(e, inf int, op func(x, y int) int, mapping func(x, y int) int, composition func(x, y int) int) {
	tree.e = e
	tree.inf = inf
	tree.op = op
	tree.mapping = mapping
	tree.composition = composition
}

func (tree *LazySegmentTree) Set(i, v int) {
	i += tree.size
	tree.node[i] = v
	for shift := 1; shift <= 60; shift++ {
		tree.update(i >> shift)
	}
}

func (tree *LazySegmentTree) Prod(l, r int) int {
	if l == r {
		return tree.e
	}
	l += tree.size
	r += tree.size

	for i := tree.log; i >= 1; i-- {
		if (l>>i)<<i != l {
			tree.push(l >> i)
		}
		if (r>>i)<<i != r {
			tree.push((r - 1) >> i)
		}
	}

	sl, sr := tree.e, tree.e
	for l < r {
		if (l & 1) > 0 {
			sl = tree.op(sl, tree.node[l])
			l++
		}
		if (r & 1) > 0 {
			r--
			sr = tree.op(sr, tree.node[r])
		}
		l >>= 1
		r >>= 1
	}
	return tree.op(sl, sr)
}

func (tree *LazySegmentTree) Apply(l, r int, f int) {
	if l == r {
		return
	}
	l += tree.size
	r += tree.size

	for i := tree.log; i >= 1; i-- {
		if (l>>i)<<i != l {
			tree.push(l >> i)
		}
		if (r>>i)<<i != r {
			tree.push((r - 1) >> i)
		}
	}

	{
		l2, r2 := l, r
		for l < r {
			if (l & 1) > 0 {
				tree.allApply(l, f)
				l++
			}
			if (r & 1) > 0 {
				r--
				tree.allApply(r, f)
			}
			l >>= 1
			r >>= 1
		}
		l, r = l2, r2
	}

	for i := 1; i <= tree.log; i++ {
		if (l>>i)<<i != l {
			tree.update(l >> i)
		}
		if (r>>i)<<i != r {
			tree.update((r - 1) >> i)
		}
	}

}

func (tree *LazySegmentTree) update(k int) {
	tree.node[k] = tree.op(tree.node[2*k], tree.node[2*k+1])
}

func (tree *LazySegmentTree) allApply(k int, f int) {
	tree.node[k] = tree.mapping(f, tree.node[k])
	if k < tree.size {
		tree.lazy[k] = tree.composition(f, tree.lazy[k])
	}
}

func (tree *LazySegmentTree) push(k int) {
	tree.allApply(2*k, tree.lazy[k])
	tree.allApply(2*k+1, tree.lazy[k])
	tree.lazy[k] = tree.inf
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
