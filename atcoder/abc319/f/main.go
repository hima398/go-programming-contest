package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/liyue201/gostl/ds/priorityqueue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var p, t, s, g []int
	for i := 0; i < n-1; i++ {
		p = append(p, nextInt()-1)
		t = append(t, nextInt())
		s = append(s, nextInt())
		g = append(g, nextInt())
	}
	if solve(n, p, t, s, g) {
		Print("Yes")
	} else {
		Print("No")
	}
}

type node struct {
	i, t, s, g int
}

func EnemyComparator(a, b node) int {
	if a.s == b.s {
		return 0
	}
	if a.s < b.s {
		return -1
	}
	return 1
}

type medicine struct {
	i, g   int
	isUsed bool
}

func solve(n int, p, t, s, g []int) bool {
	//探索するための木の構築
	//辺
	m := n - 1
	e := make([][]int, n)
	for j := 0; j < m; j++ {
		e[j+1] = append(e[j+1], p[j])
		e[p[j]] = append(e[p[j]], j+1)
	}
	//fmt.Println(e)
	//頂点
	ns := []node{node{0, 0, 0, 0}}
	for i := 1; i < n; i++ {
		ns = append(ns, node{i, t[i-1], s[i-1], g[i-1]})
	}
	//fmt.Println(ns)
	q := priorityqueue.New[node](EnemyComparator, priorityqueue.WithGoroutineSafe())
	var ms []medicine
	//q2 := priorityqueue.New[int](comparator.IntComparator, priorityqueue.WithGoroutineSafe())
	//q1 := set.NewMultiSet[node](NodeComparator, set.WithGoroutineSafe())

	strength := 1
	visited := make([]bool, n)
	visited[0] = true

	var dfs func(cur, par int)
	dfs = func(cur, par int) {
		if !visited[cur] {
			switch ns[cur].t {
			case 1:
				if strength >= ns[cur].s { //敵をその場で倒す
					strength += ns[cur].g
					visited[cur] = true
				} else { //後回しにする
					q.Push(ns[cur])
					return
				}
			case 2:
				found := false
				for _, v := range ms {
					found = found || cur == v.i
				}
				if !found {
					ms = append(ms, medicine{ns[cur].i, ns[cur].g, false})
					sort.Slice(ms, func(i, j int) bool {
						return ms[i].g < ms[j].g
					})
				}
			}
		}
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			dfs(next, cur)
		}
	}

	for i := 0; i < n; i++ {
		//fmt.Println("s = ", strength)
		dfs(0, -1)
		//fmt.Println(visited)
		//fmt.Println(*q)
		//fmt.Println(ms)
		//後回しにした敵を処理できればしておく
		for !q.Empty() {
			enemy := q.Pop()
			if strength >= enemy.s {
				strength += enemy.g
				visited[enemy.i] = true
			} else {
				q.Push(enemy)
				break
			}
		}
		nextTarget := -1
		if !q.Empty() {
			enemy := q.Pop()
			nextTarget = enemy.s
		}
		//最適な薬を使用する
		if nextTarget < 0 && len(ms) > 0 {
			for i, v := range ms {
				if v.isUsed {
					continue
				}
				strength *= ms[i].g
				ms[i].isUsed = true
				visited[ms[i].i] = true
				break
			}
		} else {
			for i, v := range ms {
				if v.isUsed {
					continue
				}
				strength *= v.g
				ms[i].isUsed = true
				visited[ms[i].i] = true
				if strength*v.g >= nextTarget {
					break
				}
			}
		}
	}
	ok := true
	for _, v := range visited {
		ok = ok && v
	}
	return ok
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
