package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

const INF = 1 << 60

type Query struct {
	h1, h2, w1, w2 int
}

func solve(n, q int, a, b []int, qs []Query) (ans []int) {
	// 0-indexed
	for k := range qs {
		qs[k].h1--
		qs[k].h2--
		qs[k].w1--
		qs[k].w2--
	}

	f := func(x1, x2 int) int { return Gcd(x1, x2) }
	as := NewSegmentTree(n-1, 0, f)
	bs := NewSegmentTree(n-1, 0, f)
	for i := 1; i < n; i++ {
		as.Update(i-1, Abs(a[i]-a[i-1]))
		bs.Update(i-1, Abs(b[i]-b[i-1]))
	}
	//fmt.Println(as.nodes)
	//fmt.Println(bs.nodes)
	for _, query := range qs {
		var gh, gw int
		v := a[query.h1] + b[query.w1]
		if query.h1 == query.h2 {
			gh = v
		} else {
			gh = as.Query(query.h1, query.h2)
		}
		if query.w1 == query.w2 {
			gw = v
		} else {
			gw = bs.Query(query.w1, query.w2)
		}
		//fmt.Println(v, gh, gw)
		v = Gcd(v, Gcd(gh, gw))
		ans = append(ans, v)
	}
	return ans
}

func solveHonestly(n, q int, a, b []int, qs []Query) (ans []int) {
	for k := range qs {
		qs[k].h1--
		qs[k].h2--
		qs[k].w1--
		qs[k].w2--
	}
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = a[i] + b[j]
		}
	}
	for k := 0; k < q; k++ {
		var gcd int
		for i := qs[k].h1; i <= qs[k].h2; i++ {
			for j := qs[k].w1; j <= qs[k].w2; j++ {
				if gcd == 0 {
					gcd = f[i][j]
				} else {
					gcd = Gcd(gcd, f[i][j])
				}
			}
		}
		ans = append(ans, gcd)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	a, b := nextIntSlice(n), nextIntSlice(n)
	var qs []Query
	for i := 0; i < q; i++ {
		qs = append(qs, Query{nextInt(), nextInt(), nextInt(), nextInt()})
	}

	//ans := solveHonestly(n, q, a, b, qs)
	ans := solve(n, q, a, b, qs)

	PrintSlice(ans)
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

func Print(a interface{}) {
	defer out.Flush()
	fmt.Fprintln(out, a)
}

func PrintSlice(a []int) {
	defer out.Flush()

	for _, v := range a {
		fmt.Fprintln(out, v)
	}
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

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	/*
		if x < y {
			x, y = y, x
		}
	*/
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	// x*yのオーバーフロー対策のため先にGcdで割る
	// Gcd(x, y)はxの約数のため割り切れる
	ret := x / Gcd(x, y)
	ret *= y
	return ret
}
