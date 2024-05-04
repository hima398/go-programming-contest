package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var nf int
var sf int64

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	a := nextIntSlice(n)
	var t []int
	p, x, l, r := make([]int, q), make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		switch t[i] {
		case 1:
			p[i], x[i] = nextInt()-1, nextInt()
		case 2:
			l[i], r[i] = nextInt()-1, nextInt()-1
		}
	}

	ans := solve(n, q, a, t, p, x, l, r)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, q int, a, t, p, x, l, r []int) []int {
	seg := NewSegmentTree(n)
	for i, ai := range a {
		seg.Update(i, ai)
	}

	var ans []int
	var n1, n2 int
	var s1, s2 int64
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			n1++
			s := time.Now().UnixNano()
			a[p[i]] = x[i]
			seg.Update(p[i], x[i])
			s1 += time.Now().UnixNano() - s
		case 2:
			//fmt.Println(seg.nodes)
			n2++
			s := time.Now().UnixNano()
			max := seg.Query(l[i], r[i]+1)
			//if max[0].v == max[1].v {
			//ans = append(ans, 0)
			//} else {
			ans = append(ans, max[1].c)
			//}
			s2 += time.Now().UnixNano() - s
		}
	}
	//fmt.Println("Average: case 1:", int(s1)/n1, ", case 2:", int(s2)/n2)
	//fmt.Println("Total: f:", int(sf), " Times: f:", nf, " Average: f:", int(sf)/nf)
	return ans
}

type node struct {
	v, c int
}

type SegmentTree struct {
	size  int
	nodes [][]node
	inf   []node
}

func NewSegmentTree(n int) *SegmentTree {
	st := new(SegmentTree)
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([][]node, 2*st.size)
	for i := range st.nodes {
		st.nodes[i] = make([]node, 2)
	}

	st.inf = []node{{0, 0}, {0, 0}}
	return st
}

func (seg *SegmentTree) f(x, y []node) []node {
	nf++
	s := time.Now().UnixNano()
	res := make([]node, len(x))
	copy(res, x)
	for _, yi := range y {
		if yi.v > res[0].v {
			res = []node{{yi.v, yi.c}, {res[0].v, res[0].c}}
		} else if yi.v == res[0].v {
			res[0].c += yi.c
		} else if yi.v > res[1].v {
			res = []node{{res[0].v, res[0].c}, {yi.v, yi.c}}
		} else if yi.v == res[1].v {
			res[1].c += yi.c
		}
	}
	sf += time.Now().UnixNano() - s
	return res
}

func (seg *SegmentTree) queryRecursively(a, b, k, l, r int) []node {
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
	return seg.f(vl, vr)
}

func (seg *SegmentTree) Query(l, r int) []node {
	return seg.queryRecursively(l, r, 1, 0, seg.size)
}

func (seg *SegmentTree) Update(k, x int) {
	k += seg.size
	seg.nodes[k] = []node{{x, 1}, {0, 0}}
	for k > 1 {
		//k /= 2
		k >>= 1
		seg.nodes[k] = seg.f(seg.nodes[k<<1], seg.nodes[k<<1+1])
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
