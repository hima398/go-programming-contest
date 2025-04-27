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

	n, m := nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(m)

	ans := solve(n, m, a, b)

	PrintHorizonaly(ans)
}

func solve(n, m int, a, b []int) []int {
	segTree := NewLazySegmentTree(2 * n)
	op := func(x, y [2]int) [2]int {
		return [2]int{x[0] + y[0], x[1] + y[1]}
	}
	mapping := func(x [2]int, y int) [2]int {
		return [2]int{x[0] + y*x[1], x[1]}
	}
	composition := func(x, y int) int {
		return x + y
	}

	segTree.Init([2]int{0, 0}, 0, op, mapping, composition)
	s := make([]int, m+1)
	c := make([]int, n)
	for k, bi := range b {
		//a[bi]からボールを取り出す
		v1 := segTree.Prod(bi, bi+1)
		v2 := segTree.Prod(bi+n, bi+n+1)
		ai := a[bi] + (s[k] - s[c[bi]]) + v1[0] + v2[0]
		//fmt.Println(k, " out = ", ai)
		//fmt.Println("segTree = ", v1, v2)
		a[bi] = 0
		c[bi] = k

		segTree.Apply(bi, bi+1, -v1[0])
		segTree.Apply(bi+n, bi+n+1, -v2[0])

		//手順に沿ってボールを配る
		s[k+1] += s[k] + (ai / n)
		t := ai % n
		idx := (bi + 1) % n
		segTree.Apply(idx, idx+t, 1)
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		v1, v2 := segTree.Prod(i, i+1), segTree.Prod(i+n, i+n+1)
		ans[i] = a[i] + (s[m] - s[c[i]]) + v1[0] + v2[0]
		//fmt.Printf("i:%d, a[i]:%d, s:%d, c[i]:%d, v1:%d, v2:%d\n", i, a[i], s, c[i], v1, v2)
	}
	return ans
}

type LazySegmentTree struct {
	n           int
	size        int
	log         int
	e           [2]int //二項演算における単位元
	inf         int
	node        [][2]int
	lazy        []int
	op          func(x, y [2]int) [2]int     //ノード間の二項演算
	mapping     func(x [2]int, y int) [2]int //ノードに遅延評価を適用する関数
	composition func(x, y int) int           //遅延評価の値xにyを合成する関数
}

func NewLazySegmentTree(n int) *LazySegmentTree {
	res := new(LazySegmentTree)
	res.n = n
	res.size = 1
	for res.size < res.n {
		res.size *= 2
	}
	res.node = make([][2]int, 2*res.size)
	for i := res.size; i < len(res.node); i++ {
		res.node[i][1] = 1
	}
	res.lazy = make([]int, 2*res.size)

	res.log = bits.TrailingZeros(uint(res.size))

	return res
}

func (tree *LazySegmentTree) Init(e [2]int, inf int, op func(x, y [2]int) [2]int, mapping func(x [2]int, y int) [2]int, composition func(x, y int) int) {
	tree.e = e
	tree.inf = inf
	tree.op = op
	tree.mapping = mapping
	tree.composition = composition
}

func (tree *LazySegmentTree) Prod(l, r int) [2]int {
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
	tree.node[k] = tree.mapping(tree.node[k], f) //木のノードkに遅延評価の値fを適用する
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
