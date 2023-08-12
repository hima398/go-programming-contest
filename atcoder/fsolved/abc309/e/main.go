package main

import (
	"bufio"
	"container/heap"
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

	n, m := nextInt(), nextInt()
	p := nextIntSlice(n - 1)
	for i := range p {
		p[i]--
	}
	var x, y []int
	for i := 0; i < m; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt())
	}
	ans := solve(n, m, p, x, y)
	PrintInt(ans)
}

func solve(n, m int, p, x, y []int) int {
	e := make([][]int, n)
	for i := range p {
		e[p[i]] = append(e[p[i]], i+1)
	}
	//頂点iが加入する一番深い保険
	g := make([]int, n)
	for i := 0; i < m; i++ {
		g[x[i]] = Max(g[x[i]], y[i])
	}

	// 点iを保証する保険
	maxStaminas := make([]int, n)
	for i := range maxStaminas {
		maxStaminas[i] = -1
	}
	q := &PriorityQueue{}
	heap.Init(q)

	push := func(to, cost int) {
		if maxStaminas[to] >= cost {
			return
		}
		maxStaminas[to] = cost
		heap.Push(q, Edge{to, cost})
	}
	for i := 0; i < n; i++ {
		if g[i] > 0 {
			push(i, g[i])
		}
		//push(p[i], y[i])
	}
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if maxStaminas[cur.t] > cur.w {
			continue
		}
		for _, next := range e[cur.t] {
			push(next, maxStaminas[cur.t]-1)
		}
	}
	var ans int
	for _, v := range maxStaminas {
		if v >= 0 {
			ans++
		}
	}
	return ans
}

func solve01(n, m int, p, x, y []int) int {
	e := make([][]int, n)
	//頂点iの深さ
	d := make([]int, n)
	for i := range p {
		e[i+1] = append(e[i+1], p[i])
		e[p[i]] = append(e[p[i]], i+1)
		d[i+1] = d[p[i]] + 1
	}
	//頂点iが加入する一番深い保険
	g := make([]int, n)
	for i := 0; i < m; i++ {
		g[x[i]] = Max(g[x[i]], y[i])
	}
	type node struct {
		i, d, g int
	}
	var ns []node
	for i := 0; i < n; i++ {
		if g[i] > 0 {
			ns = append(ns, node{i, d[i], g[i]})
		}
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].d > ns[j].d
	})
	visited := make([]bool, n)
	//uf := NewUnionFind(n)
	var dfs func(i, d, g int)
	dfs = func(i, d, g int) {
		visited[i] = true
		if d == g {
			return
		}
		for _, next := range e[i] {
			if visited[next] {
				continue
			}
			dfs(next, d+1, g)
		}
	}
	for i := range ns {
		dfs(ns[i].i, 0, ns[i].g)
	}
	var ans int
	for _, ok := range visited {
		if ok {
			ans++
		}
	}
	return ans
}

type Edge struct {
	t, w int
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].w > pq[j].w
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
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
