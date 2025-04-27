package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

	h, w, n := nextInt(), nextInt(), nextInt()
	var r, c, l []int
	for i := 0; i < n; i++ {
		r = append(r, nextInt()-1)
		c = append(c, nextInt()-1)
		l = append(l, nextInt())
	}

	ans := solve(h, w, n, r, c, l)

	for _, v := range ans {
		Print(v)
	}
}

func solve(h, w, n int, r, c, l []int) []int {
	lazySegTree := NewLazySegmentTree(w)
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	lazySegTree.Init(0, 0, max, max, max)

	type bar struct {
		idx  int
		row  int
		l, r int
	}
	var bars []bar
	for i := 0; i < n; i++ {
		bars = append(bars, bar{i, r[i], c[i], c[i] + l[i]})
	}
	sort.Slice(bars, func(i, j int) bool {
		if bars[i].row == bars[j].row {
			return bars[i].idx < bars[j].idx
		}
		return bars[i].row > bars[j].row
	})
	ans := make([]int, n)
	for _, b := range bars {
		s := lazySegTree.Prod(b.l, b.r) + 1
		ans[b.idx] = h - s + 1
		lazySegTree.Apply(b.l, b.r, s)
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
