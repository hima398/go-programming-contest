package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

	n, q := nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(n)
	var t, l, r, x []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		l, r = append(l, nextInt()), append(r, nextInt())
		if t[i] == 1 || t[i] == 2 {
			x = append(x, nextInt())
		} else {
			x = append(x, 0)
		}
	}

	ans := solve(n, q, a, b, t, l, r, x)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, q int, a, b, t, l, r, x []int) []int {
	const p = 998244353

	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = (s[i] + a[i]*b[i]%p) % p
	}

	var ans []int
	for i, ti := range t {
		switch ti {
		case 1:
		case 2:
		case 3:
			v := s[r[i]] - s[l[i]-1]
			ans = append(ans, v)
		}
	}
	return ans
}

type SegmentTree struct {
	p           int
	size        int
	nodes       []int
	lazy        []int
	op          func(a, b int) int //セグメント木のノード間の二項演算
	mapping     func(f, v int)     //セグメント木のノード値にパラメータを適用する
	composition func(g, f int)     //セグメント木の遅延評価パラメータに続けて新たなパラメータを合成する
}

func NewSegmentTree(n, p int) (st *SegmentTree) {
	st = new(SegmentTree)
	st.p = p
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([]int, 2*st.size)
	st.lazy = make([]int, 2*st.size)
	return st
}

// k番目のノードの遅延評価を行う
func (st *SegmentTree) Eval(k int) {
	if k < st.size {
		st.lazy[k*2] = (st.lazy[k*2] + st.lazy[k]) % st.p     //Max(this.lazy[k*2], this.lazy[k])
		st.lazy[k*2+1] = (st.lazy[k*2+1] + st.lazy[k]) % st.p //Max(this.lazy[k*2+1], this.lazy[k])
	}
	st.nodes[k] = Max(st.nodes[k], st.lazy[k])
	st.lazy[k] = 0
}

func (st *SegmentTree) updateRecursively(a, b, x, k, l, r int) {
	st.Eval(k)
	if a >= r || b <= l {
		return
	}
	if a <= l && b >= r {
		st.lazy[k] = x
		st.Eval(k)
		return
	}
	st.updateRecursively(a, b, x, 2*k, l, (l+r)/2)
	st.updateRecursively(a, b, x, 2*k+1, (l+r)/2, r)
	st.nodes[k] = (st.nodes[2*k] + st.nodes[2*k+1]) % st.p //Max(this.nodes[2*k], this.nodes[2*k+1])
}

func (st *SegmentTree) Update(l, r, x int) {
	st.updateRecursively(l, r, x, 1, 0, st.size)
}

func (st *SegmentTree) queryRecursively(a, b, k, l, r int) int {
	st.Eval(k)
	if a >= r || b <= l {
		return -1
	}
	if a <= l && b >= r {
		return st.nodes[k]
	}
	lv := st.queryRecursively(a, b, 2*k, l, (l+r)/2)
	rv := st.queryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return (lv + rv) % st.p // Max(lv, rv)
}

func (st *SegmentTree) Query(l, r int) int {
	return st.queryRecursively(l, r, 1, 0, st.size)
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

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
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

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

type Comb struct {
	n, p int
	fac  []int // Factional(i) mod p
	finv []int // 1/Factional(i) mod p
	inv  []int // 1/i mod p
}

func NewCombination(n, p int) *Comb {
	c := new(Comb)
	c.n = n
	c.p = p
	c.fac = make([]int, n+1)
	c.finv = make([]int, n+1)
	c.inv = make([]int, n+1)

	c.fac[0] = 1
	c.fac[1] = 1
	c.finv[0] = 1
	c.finv[1] = 1
	c.inv[1] = 1
	for i := 2; i <= n; i++ {
		c.fac[i] = c.fac[i-1] * i % p
		c.inv[i] = p - c.inv[p%i]*(p/i)%p
		c.finv[i] = c.finv[i-1] * c.inv[i] % p
	}
	return c
}

func (c *Comb) Factional(x int) int {
	return c.fac[x]
}

func (c *Comb) Combination(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	ret := c.fac[n] * c.finv[k]
	ret %= c.p
	ret *= c.finv[n-k]
	ret %= c.p
	return ret
}

// 重複組み合わせ H
func (c *Comb) DuplicateCombination(n, k int) int {
	return c.Combination(n+k-1, k)
}
func (c *Comb) Inv(x int) int {
	return c.inv[x]
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}
