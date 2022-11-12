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

type Node struct {
	i, c1, c2 int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	//if pq[i].c2 == pq[j].c2 {
	//	return pq[i].c1 > pq[i].c1
	//}
	return pq[i].c1 < pq[j].c1
	//return pq[i].c < pq[j].c
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func solveGreedily(n, m int, a, u, v []int) int {
	//e := make([][]int, n)
	e := make([]map[int]struct{}, n)
	for i := 0; i < m; i++ {
		u[i]--
		v[i]--
		if e[u[i]] == nil {
			e[u[i]] = make(map[int]struct{})
		}
		//e[u[i]] = append(e[u[i]], v[i])
		e[u[i]][v[i]] = struct{}{}
		if e[v[i]] == nil {
			e[v[i]] = make(map[int]struct{})
		}
		//e[v[i]] = append(e[v[i]], u[i])
		e[v[i]][u[i]] = struct{}{}
	}
	//点iを削除する最小のコスト
	sa := make([]int, n)
	for i := 0; i < n; i++ {
		for next := range e[i] {
			sa[i] += a[next]
		}
	}
	q := &PriorityQueue{}
	heap.Init(q)
	for i := range sa {
		heap.Push(q, Node{i, sa[i], 0})
	}
	deleted := make([]bool, n)
	var ans int
	for q.Len() > 0 {
		cur := heap.Pop(q).(Node)
		if deleted[cur.i] {
			continue
		}
		deleted[cur.i] = true
		ans = Max(ans, cur.c1)
		for next := range e[cur.i] {
			if deleted[next] {
				continue
			}
			sa[next] -= a[cur.i]
			heap.Push(q, Node{next, sa[next], 0})
		}
	}

	return ans
}

func solve(n, m int, a, u, v []int) int {
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		u[i]--
		v[i]--
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	mx := 0
	for _, ai := range a {
		mx += ai
	}
	s := make([]int, n)
	//O(n+m)
	for i := 0; i < n; i++ {
		for _, next := range e[i] {
			s[i] += a[next]
		}
	}
	ng, ok := -1, mx
	check := func(x int) bool {
		if m == 0 {
			return true
		}
		var ns []Node
		deleted := make([]bool, n)
		ss := make([]int, n)
		copy(ss, s)
		for i := 0; i < n; i++ {
			if ss[i] <= x {
				ns = append(ns, Node{i, a[i], ss[i]})
				deleted[i] = true
			}
		}
		for len(ns) > 0 {
			node := ns[0]
			ns = ns[1:]
			for _, next := range e[node.i] {
				if deleted[next] {
					continue
				}
				ss[next] -= a[node.i]
				if ss[next] <= x {
					ns = append(ns, Node{next, a[next], ss[next]})
					deleted[next] = true
				}
			}
		}
		ok := true
		for _, b := range deleted {
			ok = ok && b
		}
		return ok
	}
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := nextIntSlice(n)
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt())
		v = append(v, nextInt())
	}
	//ans := solve(n, m, a, u, v)
	ans := solveGreedily(n, m, a, u, v)
	PrintInt(ans)
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
