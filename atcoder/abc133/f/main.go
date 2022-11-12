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

func solve(n, q int, a, b, c, d, x, y, u, v []int) (ans []int) {
	type Query struct {
		i, x, y, w int
	}
	queue := make([][]Query, n)
	tree := NewTree(n, n-1, a, b)
	tree.InitLca(n)
	for i := 0; i < q; i++ {
		queue[u[i]] = append(queue[u[i]], Query{i, x[i], y[i], 1})
		queue[v[i]] = append(queue[v[i]], Query{i, x[i], y[i], 1})
		par := tree.Lca(u[i], v[i])
		queue[par] = append(queue[par], Query{i, x[i], y[i], -2})
	}
	var dfs func(cur, par, s int)
	dfs = func(cur, par, s int) {
		for _, next := range tree.e[cur] {
			if next == par {
				continue
			}
		}
	}
	dfs(0, -1, 0)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	var a, b, c, d []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt()-1)
		d = append(d, nextInt())
	}
	var x, y, u, v []int
	for i := 0; i < q; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt())
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, q, a, b, c, d, x, y, u, v)
	PrintVertically(ans)
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

//重複組み合わせ H
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
