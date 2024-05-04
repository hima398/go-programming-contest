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

	n, d := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, d, a)
	Print(ans)
}

func solve(n, d int, a []int) int {
	var mx int
	for _, ai := range a {
		mx = Max(mx, ai)
	}
	dp := NewSegmentTree(mx+1, 0, func(x, y int) int {
		if x < y {
			return y
		}
		return x
	})
	for _, ai := range a {
		dp.Update(ai, dp.Query(Max(0, ai-d), Min(mx, ai+d)+1)+1)
	}
	return dp.nodes[1]
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

func (seg *SegmentTree) Query(l, r int) int {
	return seg.QueryRecursively(l, r, 1, 0, seg.size)
}

func (seg *SegmentTree) Update(k, x int) {
	k += seg.size
	seg.nodes[k] = seg.f(seg.nodes[k], x)
	for k > 1 {
		k /= 2
		seg.nodes[k] = seg.f(seg.nodes[k*2], seg.nodes[k*2+1])
	}
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
