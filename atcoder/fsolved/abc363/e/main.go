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
	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, y := nextInt(), nextInt(), nextInt()
	var a [][]int
	for i := 0; i < h; i++ {
		a = append(a, nextIntSlice(w))
		for j := range a[i] {
			a[i][j]--
		}
	}
	ans := solve(h, w, y, a)
	for _, v := range ans {
		Print(v)
	}

}

func solve(h, w, y int, a [][]int) []int {
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}

	t := make([]int, y+1)
	type node struct {
		h    int
		i, j int
	}
	sinked := make([][]bool, h)
	evented := make([][]bool, h)
	for i := 0; i < h; i++ {
		sinked[i] = make([]bool, w)
		evented[i] = make([]bool, w)
	}
	pq := priorityqueue.New[node](func(a, b node) int {
		if a.h == b.h {
			return 0
		}
		if a.h < b.h {
			return -1
		}
		return 1
	})
	for i := 0; i < h; i++ {
		pq.Push(node{a[i][0], i, 0})
		pq.Push(node{a[i][w-1], i, w - 1})
		evented[i][0] = true
		evented[i][w-1] = true
	}
	for j := 0; j < w; j++ {
		if j == 0 || j == w-1 {
			continue
		}
		pq.Push(node{a[0][j], 0, j})
		pq.Push(node{a[h-1][j], h - 1, j})
		evented[0][j] = true
		evented[h-1][j] = true
	}
	type pos struct {
		i, j int
	}
	for !pq.Empty() {
		sink := pq.Pop()
		if sink.h > y {
			break
		}
		if sinked[sink.i][sink.j] {
			continue
		}
		q := queue.New[pos]()
		q.Push(pos{sink.i, sink.j})
		sinked[sink.i][sink.j] = true
		cnt := 1
		for !q.Empty() {
			cur := q.Pop()
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				if sinked[ni][nj] {
					continue
				}
				if a[ni][nj] > sink.h {
					if !evented[ni][nj] && a[ni][nj] <= y {
						pq.Push(node{a[ni][nj], ni, nj})
						evented[ni][nj] = true
					}
					continue
				}
				q.Push(pos{ni, nj})
				sinked[ni][nj] = true
				cnt++
			}
		}
		//沈んだ土地を計算する
		t[sink.h] -= cnt
	}
	//fmt.Println(t)
	for i := 1; i <= y; i++ {
		t[i] += t[i-1]
	}
	ans := make([]int, y)
	for i := range ans {
		ans[i] = h*w + t[i]
	}
	return ans
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
	/*
		return x / y
	*/
	r := (x%y + y) % y
	return (x - r) / y
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

// xの平方根の整数部分を計算する
func FloorSqrt(x int) int {
	l, r := 0, int(2e9)
	for l <= r {
		mid := (l + r) / 2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
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
