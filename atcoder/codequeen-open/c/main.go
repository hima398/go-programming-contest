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

	n, s, t := nextInt(), nextInt()-1, nextInt()-1
	var u, v []int
	for i := 0; i < n-1; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, s, t, u, v)
	PrintVertically(ans)
}

func solve(n, s, t int, u, v []int) []int {
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}

	const INF = 1 << 60
	d := make([]int, n)
	for i := range d {
		d[i] = INF
	}
	bfs1 := func(s int) {
		var q []int
		q = append(q, s)
		d[s] = 1
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				if d[next] != INF {
					continue
				}
				q = append(q, next)
				d[next] = d[cur] + 1
			}
		}
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = INF
	}
	visited := make([]bool, n)
	bfs3 := func(t int) {
		var q []int
		q = append(q, t)
		visited[t] = true
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			if cur == s {
				return
			}
			for _, next := range e[cur] {
				if d[next] >= d[cur] {
					continue
				}
				q = append(q, next)
				visited[next] = true
			}
		}
	}
	bfs2 := func(s int) {
		q := &PriorityQueue{}
		push := func(to, cost int) {
			if ans[to] < cost {
				return
			} else if ans[to] > cost {
				ans[to] = cost
				heap.Push(q, Edge{to, cost})
			} else {
				//ans[to]==cost
				ans[to] = 1
				heap.Push(q, Edge{to, 1})
			}
		}

		heap.Init(q)
		for i := range visited {
			if visited[i] {
				push(i, 1)
			}
		}
		for q.Len() > 0 {
			cur := heap.Pop(q).(Edge)
			for _, next := range e[cur.t] {
				push(next, ans[cur.t]+1)
			}
		}
	}
	bfs1(s)
	bfs3(t)
	bfs2(t)
	return ans
}

func firstsolve(n, s, t int, u, v []int) []int {
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	tree := NewTree(n, n-1, u, v)
	tree.InitLca(n)

	root := tree.Lca(s, t)
	fmt.Println("root = ", root)

	q := &PriorityQueue{}
	heap.Init(q)

	const INF = 1 << 60
	ans := make([]int, n)
	for i := range ans {
		ans[i] = INF
	}
	push := func(to, cost int) {
		if ans[to] < cost {
			return
		}
		ans[to] = cost
		heap.Push(q, Edge{to, cost})
	}

	push(s, 1)
	push(t, 1)
	push(root, 1)

	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if ans[cur.t] > cur.w {
			continue
		}
		for _, next := range e[cur.t] {
			push(next, ans[cur.t]+1)
		}
	}
	return ans

}

type Tree struct {
	n, m   int
	e      [][]int
	depth  []int
	parent [][]int
}

func New(n, m int, u, v []int) *Tree {
	return NewTree(n, m, u, v)
}

func NewTree(n, m int, u, v []int) *Tree {
	tree := new(Tree)
	tree.n = n
	tree.m = m
	tree.e = make([][]int, n)
	for i := 0; i < m; i++ {
		tree.e[u[i]] = append(tree.e[u[i]], v[i])
		tree.e[v[i]] = append(tree.e[v[i]], u[i])
	}
	tree.depth = make([]int, n+1)
	tree.parent = make([][]int, 18)
	for i := 0; i <= 17; i++ {
		tree.parent[i] = make([]int, n+1)
	}
	return tree
}

func (tree *Tree) InitLca(v int) {
	for i := 0; i < 17; i++ {
		for j := 0; j < v; j++ {
			if tree.parent[i][j] < 0 {
				tree.parent[i+1][j] = -1
			} else {
				tree.parent[i+1][j] = tree.parent[i][tree.parent[i][j]]
			}
		}
	}
}

func (tree *Tree) Lca(u, v int) int {
	// depth(u) <= depth(v)になるように調整する
	if tree.depth[u] > tree.depth[v] {
		u, v = v, u
	}
	for i := 0; i < 18; i++ {
		if ((tree.depth[v]-tree.depth[u])>>i)&1 == 1 {
			v = tree.parent[i][v]
		}
	}
	if u == v {
		return u
	}
	for i := 17; i >= 0; i-- {
		if tree.parent[i][u] != tree.parent[i][v] {
			u = tree.parent[i][u]
			v = tree.parent[i][v]
		}
	}
	return tree.parent[0][u]
}

type Edge struct {
	t, w int
}
type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].w < pq[j].w
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
