package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var startTime int64

func main() {
	startTime = time.Now().UnixNano()
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, k := nextInt(), nextInt(), nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	var u, v, w []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		w = append(w, nextInt())
	}
	var a, b []int
	for i := 0; i < k; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	//p, c := solveHonestly(n, m, k, u, v, w, a, b)
	//p, c := solve01(n, m, k, x, y, u, v, w, a, b)
	//p, c := solve02(n, m, k, x, y, u, v, w, a, b)
	//p, c := solve03(n, m, k, x, y, u, v, w, a, b)
	//p, c := solve04(n, m, k, x, y, u, v, w, a, b)
	p, c := solve05(n, m, k, x, y, u, v, w, a, b)

	PrintHorizonaly(p)
	PrintHorizonaly(c)
}

var uf *UnionFind

func computeDist2(x1, y1, x2, y2 int) int {
	xx := (x2 - x1)
	yy := (y2 - y1)
	return xx*xx + yy*yy
}

func computeScore(n, m, k int, x, y, w, a, b, p, c []int) (int, bool) {
	isBroadcasted := make([]bool, k)
	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			if !uf.ExistSameUnion(0, j) {
				continue
			}
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist <= p[j]*p[j] {
				isBroadcasted[i] = true
			}
		}
	}
	var sn int
	for _, ok := range isBroadcasted {
		if ok {
			sn++
		}
	}
	if sn == k {
		var s int
		for i := 0; i < n; i++ {
			s += p[i] * p[i]
		}
		for i := 0; i < m; i++ {
			if c[i] == 0 {
				continue
			}
			s += w[i]
		}
		return int(1e6 * (1.0 + 1e8/float64(s+1e7))), true
	} else {
		return int(1e6 * float64(sn+1.0) / float64(k)), false
	}
}

// コストが少ないケーブルを使った後、一番近い放送局から受信する
// 効率の悪い出力を全探索的に改善していく
func solve05(n, m, k int, x, y, u, v, w, a, b []int) (p, c []int) {

	uf = NewUnionFind(n)

	e := make([][]Edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], Edge{i, u[i], v[i], w[i]})
		e[v[i]] = append(e[v[i]], Edge{i, v[i], u[i], w[i]})
	}

	c = make([]int, m)
	ed := make([][]int, n)
	for i := range ed {
		ed[i] = make([]int, n)
		for j := range ed[i] {
			ed[i][j] = -1
		}
	}
	q := &PriorityQueue{}
	heap.Init(q)
	for i := 0; i < m; i++ {
		heap.Push(q, Edge{i, u[i], v[i], w[i]})
	}
	for uf.Size(0) < n {
		node := heap.Pop(q).(Edge)
		if uf.ExistSameUnion(node.s, node.t) {
			continue
		}
		uf.Unite(node.s, node.t)
		c[node.i] = 1
		ed[node.s][node.t] = node.i
		ed[node.t][node.s] = node.i
	}

	p = make([]int, n)
	for i := 0; i < k; i++ {
		var isCovered bool
		for j := 0; j < n; j++ {
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist < p[j]*p[j] {
				isCovered = true
			}
		}
		if isCovered {
			continue
		}
		power := 5000
		idx := -1
		for j := 0; j < n; j++ {
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist < power*power {
				power = Min(Sqrt(dist)+1, 5000)
				idx = j
			}
		}
		p[idx] = Min(Max(p[idx], power), 5000)
	}

	//for time.Now().UnixNano()-startTime <= int64(1.7*1e9) {
	for idx := 0; idx < n; idx++ {
		np := make([]int, n)
		copy(np, p)
		if np[idx] == 0 {
			continue
		}
		np[idx] = 0
		for i := 0; i < k; i++ {
			var isCovered bool
			for j := 0; j < n; j++ {
				dist := computeDist2(a[i], b[i], x[j], y[j])
				if dist < np[j]*np[j] {
					isCovered = true
				}
			}
			if isCovered {
				continue
			}
			power := 5000
			idx := -1
			for j := 0; j < n; j++ {
				if j == idx {
					continue
				}
				dist := computeDist2(a[i], b[i], x[j], y[j])
				if dist < power*power {
					power = Min(Sqrt(dist)+1, 5000)
					idx = j
				}
			}
			np[idx] = Min(Max(np[idx], power), 5000)
		}
		score, _ := computeScore(n, m, k, x, y, w, a, b, p, c)
		newScore, newOk := computeScore(n, m, k, x, y, w, a, b, np, c)
		//fmt.Println(newScore, newOk, score)
		//住人全員をカバーできていなければ不採用
		if !newOk {
			continue
		}
		if newScore > score {
			p = np
			//fmt.Println("swaped")
		}
	}

	for time.Now().UnixNano()-startTime <= int64(1.8*1e9) {
		np := make([]int, n)
		copy(np, p)
		idx := rand.Intn(len(np))
		if p[idx] == 0 {
			continue
		}
		np[idx] = rand.Intn(p[idx])
		for i := 0; i < k; i++ {
			var isCovered bool
			for j := 0; j < n; j++ {
				dist := computeDist2(a[i], b[i], x[j], y[j])
				if dist < np[j]*np[j] {
					isCovered = true
				}
			}
			if isCovered {
				continue
			}
			power := 5000
			idx := -1
			for j := 0; j < n; j++ {
				if j == idx {
					continue
				}
				dist := computeDist2(a[i], b[i], x[j], y[j])
				if dist < power*power {
					power = Min(Sqrt(dist)+1, 5000)
					idx = j
				}
			}
			np[idx] = Min(Max(np[idx], power), 5000)
		}
		score, _ := computeScore(n, m, k, x, y, w, a, b, p, c)
		newScore, newOk := computeScore(n, m, k, x, y, w, a, b, np, c)
		//fmt.Println(newScore, newOk, score)
		//住人全員をカバーできていなければ不採用
		if !newOk {
			continue
		}
		if newScore > score {
			p = np
			//fmt.Println("swaped")
		}
	}

	//不要なケーブルの掃除
	for i := 0; i < n; i++ {
		if p[i] > 0 {
			continue
		}
		//出力がない放送局
		var numDegree int
		for j := range ed[i] {
			if ed[i][j] >= 0 {
				numDegree++
			}
		}
		if numDegree > 1 {
			continue
		}
		for j := range ed[i] {
			if ed[i][j] >= 0 {
				c[ed[i][j]] = 0
				ed[j][i] = -1
				ed[i][j] = -1
			}
		}
	}

	return p, c
}

// コストが少ないケーブルを使った後、一番近い放送局から受信する
// 効率の悪い出力を探索的に改善していく
func solve04(n, m, k int, x, y, u, v, w, a, b []int) (p, c []int) {

	uf = NewUnionFind(n)

	e := make([][]Edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], Edge{i, u[i], v[i], w[i]})
		e[v[i]] = append(e[v[i]], Edge{i, v[i], u[i], w[i]})
	}

	c = make([]int, m)
	q := &PriorityQueue{}
	heap.Init(q)
	for i := 0; i < m; i++ {
		heap.Push(q, Edge{i, u[i], v[i], w[i]})
	}
	for uf.Size(0) < n {
		node := heap.Pop(q).(Edge)
		if uf.ExistSameUnion(node.s, node.t) {
			continue
		}
		uf.Unite(node.s, node.t)
		c[node.i] = 1
	}

	p = make([]int, n)
	for i := 0; i < k; i++ {
		var isCovered bool
		for j := 0; j < n; j++ {
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist < p[j]*p[j] {
				isCovered = true
			}
		}
		if isCovered {
			continue
		}
		power := 5000
		idx := -1
		for j := 0; j < n; j++ {
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist < power*power {
				power = Min(Sqrt(dist)+1, 5000)
				idx = j
			}
		}
		p[idx] = Min(Max(p[idx], power), 5000)
	}

	for time.Now().UnixNano()-startTime <= int64(1.7*1e9) {
		np := make([]int, n)
		copy(np, p)
		idx := rand.Intn(len(np))
		if np[idx] == 0 {
			continue
		}
		//ランダムに電波の出力を0にする
		np[idx] = 0
		for i := 0; i < k; i++ {
			var isCovered bool
			for j := 0; j < n; j++ {
				dist := computeDist2(a[i], b[i], x[j], y[j])
				if dist < np[j]*np[j] {
					isCovered = true
				}
			}
			if isCovered {
				continue
			}
			power := 5000
			idx := -1
			for j := 0; j < n; j++ {
				if j == idx {
					continue
				}
				dist := computeDist2(a[i], b[i], x[j], y[j])
				if dist < power*power {
					power = Min(Sqrt(dist)+1, 5000)
					idx = j
				}
			}
			np[idx] = Min(Max(np[idx], power), 5000)
		}
		score, _ := computeScore(n, m, k, x, y, w, a, b, p, c)
		newScore, newOk := computeScore(n, m, k, x, y, w, a, b, np, c)
		//fmt.Println(newScore, newOk, score)
		//住人全員をカバーできていなければ不採用
		if !newOk {
			continue
		}
		if newScore > score {
			p = np
			//fmt.Println("swaped")
		}
	}

	return p, c
}

// コストが少ないケーブルを使った後、一番近い放送局から受信する
func solve03(n, m, k int, x, y, u, v, w, a, b []int) (p, c []int) {
	const INF = 1 << 60

	uf = NewUnionFind(n)

	e := make([][]Edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], Edge{i, u[i], v[i], w[i]})
		e[v[i]] = append(e[v[i]], Edge{i, v[i], u[i], w[i]})
	}

	c = make([]int, m)
	q := &PriorityQueue{}
	heap.Init(q)
	for i := 0; i < m; i++ {
		heap.Push(q, Edge{i, u[i], v[i], w[i]})
	}
	for uf.Size(0) < n {
		node := heap.Pop(q).(Edge)
		if uf.ExistSameUnion(node.s, node.t) {
			continue
		}
		uf.Unite(node.s, node.t)
		c[node.i] = 1
	}

	p = make([]int, n)
	for i := 0; i < k; i++ {
		var isCovered bool
		for j := 0; j < n; j++ {
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist < p[j]*p[j] {
				isCovered = true
			}
		}
		if isCovered {
			continue
		}
		power := 5000
		idx := -1
		for j := 0; j < n; j++ {
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist < power*power {
				power = Min(Sqrt(dist)+1, 5000)
				idx = j
			}
		}
		p[idx] = Min(Max(p[idx], power), 5000)
	}

	return p, c
}

// ダイクストラで放送局を連結した後、一番近い放送局から受信する
func solve02(n, m, k int, x, y, u, v, w, a, b []int) (p, c []int) {
	const INF = 1 << 60

	uf = NewUnionFind(n)

	e := make([][]Edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], Edge{i, u[i], v[i], w[i]})
		e[v[i]] = append(e[v[i]], Edge{i, v[i], u[i], w[i]})
	}

	c = make([]int, m)

	//Dijkstra
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	q := &PriorityQueue{}
	heap.Init(q)

	push := func(i, s, t, cost int) {
		if dist[t] <= cost {
			return
		}
		dist[t] = cost
		heap.Push(q, Edge{i, s, t, cost})
	}
	for _, next := range e[0] {
		push(next.i, next.s, next.t, next.w)
	}
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if dist[cur.t] < cur.w {
			continue
		}
		if uf.ExistSameUnion(cur.s, cur.t) {
			continue
		}
		uf.Unite(cur.s, cur.t)
		c[cur.i] = 1
		for _, next := range e[cur.t] {
			push(next.i, next.s, next.t, cur.w+next.w)
		}
	}

	p = make([]int, n)
	for i := 0; i < k; i++ {
		var isCovered bool
		for j := 0; j < n; j++ {
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist < p[j]*p[j] {
				isCovered = true
			}
		}
		if isCovered {
			continue
		}
		power := 5000
		idx := -1
		for j := 0; j < n; j++ {
			dist := computeDist2(a[i], b[i], x[j], y[j])
			if dist < power*power {
				power = Min(Sqrt(dist)+1, 5000)
				idx = j
			}
		}
		p[idx] = Min(Max(p[idx], power), 5000)
	}

	return p, c
}

// ダイクストラで放送局を連結した後、放送局間の距離の最大値を出力
func solve01(n, m, k int, x, y, u, v, w, a, b []int) (p, c []int) {
	const INF = 1 << 60

	uf = NewUnionFind(n)

	e := make([][]Edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], Edge{i, u[i], v[i], w[i]})
		e[v[i]] = append(e[v[i]], Edge{i, v[i], u[i], w[i]})
	}

	c = make([]int, m)

	//Dijkstra
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	q := &PriorityQueue{}
	heap.Init(q)

	push := func(i, s, t, cost int) {
		if dist[t] <= cost {
			return
		}
		dist[t] = cost
		heap.Push(q, Edge{i, s, t, cost})
	}
	for _, next := range e[0] {
		push(next.i, next.s, next.t, next.w)
	}
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if dist[cur.t] < cur.w {
			continue
		}
		if uf.ExistSameUnion(cur.s, cur.t) {
			continue
		}
		uf.Unite(cur.s, cur.t)
		c[cur.i] = 1
		for _, next := range e[cur.t] {
			push(next.i, next.s, next.t, cur.w+next.w)
		}
	}

	p = make([]int, n)
	for i := 0; i < m; i++ {
		if c[i] == 0 {
			continue
		}
		idx1, idx2 := u[i], v[i]
		power := Sqrt(computeDist2(x[idx1], y[idx1], x[idx2], y[idx2]))
		p[idx1] = Min(Max(p[idx1], power), 5000)
		p[idx2] = Min(Max(p[idx2], power), 5000)
	}

	return p, c
}

type Edge struct {
	i, s, t, w int
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

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
}

func New(n int) *UnionFind {
	return NewUnionFind(n)
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.parent = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.parent[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	return u
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] == x {
		return x
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4
		uf.parent[x] = uf.Find(uf.parent[x])
		return uf.parent[x]
	}
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Find(x)]
}

func (uf *UnionFind) ExistSameUnion(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	// rank
	if uf.rank[x] < uf.rank[y] {
		//yがrootの木にxがrootの木を結合する
		uf.parent[x] = y
		uf.size[y] += uf.size[x]
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.parent)
	fmt.Println(u.rank)
	fmt.Println(u.size)
}

func solveHonestly(n, m, k int, u, v, w, a, b []int) (p, c []int) {
	for i := 0; i < n; i++ {
		p = append(p, 5000)
	}
	for i := 0; i < m; i++ {
		c = append(c, 1)
	}
	return p, c
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
