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

type answer struct {
	d, s int
	err  error
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	q := nextInt()
	var u, v []int
	for i := 0; i < q; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, a, s, q, u, v)
	PrintVertically(ans)
}

func solve(n int, a []int, s []string, q int, u, v []int) []answer {
	const INF = 1 << 60
	e := make([][]Edge, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == 'Y' {
				e[i] = append(e[i], Edge{j, 1, 0})
			}
		}
	}
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			//if i == j {
			//	continue
			//}
			dist[i][j] = INF
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == 'Y' {
				dist[i][j] = 1
			}
		}
	}

	souvenir := make([][]int, n)
	for i := range souvenir {
		souvenir[i] = make([]int, n)
	}
	/*
		bfs := func(x int) {
			var q []int
			visited := make([]bool, n)
			q = append(q, x)
			visited[x] = true
			dist[x][x] = 0
			souvenir[x][x] = a[x]
			for len(q) > 0 {
				cur := q[0]
				q = q[1:]
				for _, next := range e[cur] {
					if visited[next] {
						continue
					}
					q = append(q, next)
					dist[x][next] = Min(dist[x][next], dist[x][cur]+1)
					souvenir[x][next] = Max(souvenir[x][next], souvenir[x][cur]+a[next])
					visited[next] = true
				}
			}
		}
	*/
	dijkstra := func(x int) ([]int, []int) {
		dist := make([]int, n)
		s := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = INF
		}
		q := &PriorityQueue{}
		heap.Init(q)

		push := func(to, cost, ss int) {
			if dist[to] <= cost {
				return
			}
			dist[to] = cost
			s[to] = Max(s[to], ss)
			heap.Push(q, Edge{to, cost, ss})
		}
		push(x, 0, a[x])
		for q.Len() > 0 {
			cur := heap.Pop(q).(Edge)
			if dist[cur.t] < cur.w {
				continue
			}
			for _, next := range e[cur.t] {
				push(next.t, cur.w+next.w, cur.s+a[next.t])
			}
		}
		return dist, s
	}

	for i := 0; i < n; i++ {
		//bfs(i)
		d, s := dijkstra(i)
		for j := 0; j < n; j++ {
			dist[i][j] = d[j]
			souvenir[i][j] = s[j]
		}
	}
	var ans []answer
	for i := 0; i < q; i++ {
		if dist[u[i]][v[i]] == INF {
			ans = append(ans, answer{-1, -1, errors.New("Impossible")})
		} else {
			ans = append(ans, answer{dist[u[i]][v[i]], souvenir[u[i]][v[i]], nil})
		}
	}
	return ans
}

type Edge struct {
	t, w, s int
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].w == pq[j].w {
		return pq[i].s > pq[j].s
	}
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

func PrintVertically(x []answer) {
	defer out.Flush()
	for _, v := range x {
		if v.err != nil {
			fmt.Fprintln(out, "Impossible")
		} else {
			fmt.Fprintln(out, v.d, v.s)
		}
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
