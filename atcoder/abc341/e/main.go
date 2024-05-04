package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	s := nextString()
	var t, l, r []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		l = append(l, nextInt())
		r = append(r, nextInt())
	}
	ans := solve(n, q, s, t, l, r)
	for _, ok := range ans {
		if ok {
			Print("Yes")
		} else {
			Print("No")
		}
	}
}

func solve(n, q int, s string, t, l, r []int) []bool {
	seg := NewSegmentTree(n+1, 0, func(x1, x2 int) int { return x1 + x2 })
	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			seg.Update(i+1, 1)
		}
	}
	var ans []bool
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			seg.Update(l[i]-1, seg.Get(l[i]-1)^1)
			seg.Update(r[i], seg.Get(r[i])^1)
		case 2:
			ans = append(ans, seg.Query(l[i], r[i]) == r[i]-l[i])
		}
	}
	return ans
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

func (seg *SegmentTree) queryRecursively(a, b, k, l, r int) int {
	// [a, b)と[l, r)が交差しない
	if a >= r || b <= l {
		return seg.inf
	}

	// [a, b)が[l, r)を完全に含んでいる
	if a <= l && b >= r {
		return seg.nodes[k]
	}

	vl := seg.queryRecursively(a, b, 2*k, l, (l+r)/2)
	vr := seg.queryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return seg.f(vl, vr) //Max(vl, vr)
}

func (seg *SegmentTree) Get(k int) int {
	return seg.nodes[k+seg.size]
}

func (seg *SegmentTree) Query(l, r int) int {
	return seg.queryRecursively(l, r, 1, 0, seg.size)
}

func (seg *SegmentTree) Update(k, x int) {
	k += seg.size
	seg.nodes[k] = x //seg.f(seg.nodes[k], x)
	for k > 1 {
		k /= 2
		seg.nodes[k] = seg.f(seg.nodes[k*2], seg.nodes[k*2+1])
	}
}

func (seg *SegmentTree) String() string {
	var res []string
	for _, v := range seg.nodes {
		res = append(res, strconv.Itoa(v))
	}
	return strings.Join(res, " ")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
